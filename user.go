package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// User Class init function
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: fmt.Sprintf("客户%s", userAddr[len(userAddr)-5:]),
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// Starting a goroutine to listenning user channel
	go user.ListenMessage()

	return user
}

// Listening User Channel, Once have message, Transmit to Connection
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
