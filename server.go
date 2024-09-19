package main

import (
	"fmt"
	"net"
)

// define a server class
type Server struct {
	Ip   string
	Port int
}

// create a server class init function
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	// do some thing...
	fmt.Printf("listening...start a goroutine: %v\n", conn)
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
