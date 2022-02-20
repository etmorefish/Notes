package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听Message 广播消息 channel 的 goroutine， 一旦有消息就发送给在线的用户
func (this *Server) ListenMessager() {
	fmt.Println("// ListenMessager ...")
	for {
		msg := <-this.Message
		fmt.Println("// sendMsg to onlineMap")
		//将msg 发送给全部在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			fmt.Println("// range OnlineMap ...")
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	fmt.Println("// BroadCast ...")
	sendMsg := "[" + user.Name + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("// do Handler ...")
	// ...当前链接的业务
	// fmt.Println("连接建立成功")

	user := NewUser(conn, this)

	user.Online()

	//监听用户时候活跃的 channel
	isLive := make(chan bool)

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil {
				fmt.Println("Conn Read err:", err)
				return
			}

			//提取用户消息（去掉‘\n’）
			msg := string(buf[:n-1])

			//用户针对msg进行消息处理
			user.DoMessage(msg)

			//用户的任意消息 代表当前用户活跃
			isLive <- true
		}
	}()
	//当前 handler阻塞
	for {
		select {
		// 定时器
		case <-isLive:
		case <-time.After(time.Second * 20):
			//已经超时

			//将当前User 强制关闭
			user.SendMsg("您已被强制下线")

			//销毁用户资源
			close(user.C)

			//关闭连接
			conn.Close()

			//退出当前Hander
			// runtime.Goexit()
			return
		}
	}

}

//启动服务器的接口
func (this *Server) Start() {
	// TODO: socket listen
	fmt.Println("// Server starting ...")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// TODO: close socket listen
	defer listener.Close()
	defer fmt.Println("// close socket listen")

	//启动监听 Message 的 goroutine
	go this.ListenMessager()

	for {
		// TODO: accept socket
		fmt.Println("// accept socket ...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		// TODO: do handler
		go this.Handler(conn)
	}
}
