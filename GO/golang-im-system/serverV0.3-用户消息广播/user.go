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

//创建一个用户的API
func NewUser(conn net.Conn) *User {
	fmt.Println("// Create User API...")
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
	}

	//自动监听当前 user channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

//监听当前channel 的方法中一旦有消息， 就直接发生送给客户端
func (this *User) ListenMessage(){
	fmt.Println("// ListenMessage")
	for {
		msg := <-this.C
		fmt.Println("// sendMsg to client")
		this.conn.Write([]byte(msg + "\n"))
	}
}