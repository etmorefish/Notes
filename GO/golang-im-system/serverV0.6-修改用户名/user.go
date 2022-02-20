package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

//创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	fmt.Println("// Create User API...")
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//自动监听当前 user channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

//用户的上线
func (this *User) Online() {
	//用户上线，将用户加入到 onlineMap 中
	this.server.mapLock.Lock()
	fmt.Println("// new user, and add to map")
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	fmt.Println("// BroadCast user Online ...")
	this.server.BroadCast(this, "已上线")
}

//用户的下线
func (this *User) Offline() {
	//用户下线，将用户从 onlineMap 中删除
	this.server.mapLock.Lock()
	fmt.Println("//  user, and del to map")
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	fmt.Println("// BroadCast user Offline ...")
	this.server.BroadCast(this, "下线")
}

//给当前User 对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//用户的处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式: rename|Phil
		newName := strings.Split(msg, "|")[1]

		//判断name 是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名称已被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已更新用户名:" + this.Name + "\n")
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}

//监听当前channel 的方法中一旦有消息， 就直接发生送给客户端
func (this *User) ListenMessage() {
	fmt.Println("// ListenMessage")
	for {
		msg := <-this.C
		fmt.Println("// sendMsg to client")
		this.conn.Write([]byte(msg + "\n"))
	}
}
