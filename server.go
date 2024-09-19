package main

import (
	"fmt"
	"net"
	"sync"
)

// define a server class
type Server struct {
	Ip   string
	Port int

	// Online Users MapList
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// broadcast channel for message
	MessageChannel chan string
}

// create a server class init function
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:             ip,
		Port:           port,
		OnlineMap:      make(map[string]*User),
		MessageChannel: make(chan string),
	}

	return server
}

// Send Message To BroadCast Channel
func (this *Server) BroadCast(user *User, msg string) {
	sendMessage := "[" + user.Name + "]" + user.Addr + ": " + msg

	this.MessageChannel <- sendMessage
}

// Broadcasting Message To All Users
func (this *Server) MessageListener() {
	// Listenning Channel
	for {
		msg := <-this.MessageChannel

		// sending msg to all online users (COMMON MATERIALS LOCK!)
		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	// TODO: handle transaction...
	user := NewUser(conn)

	// Adding UserInfo To OnlienMap (BE CAREFUL ADD LOCK!)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// Broadcasting Online Info
	this.BroadCast(user, "已上线! ")

	// BLOCKING
	select {}
}

// open server
func (this *Server) Start() {
	// socket listen
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("network listening error: ", err)
		return
	}

	// close listen socket
	defer Listener.Close()

	// create Listener for msgChannel(JUST ONE)
	go this.MessageListener()

	// 主线程一直循环保持监听状态，当有连接请求时开启goroutine处理业务
	for {
		// Accept
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}
}
