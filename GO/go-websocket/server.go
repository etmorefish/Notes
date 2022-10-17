package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		message = []byte(string(message) + "modfity")
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

/*
Gorilla的工作是转换原始HTTP连接进入一个有状态的websocket连接。
这就是为什么使用struct调用Upgrader来帮助我们的原因。
我们使用全局升级程序变量通过来帮助我们将任何传入的HTTP连接转换为websocket协议upgrader.Upgrade()。
这将返回给我们*websocket.Connection，我们现在可以使用它来处理websocket连接。
服务器使用读取消息，然后使用conn.ReadMessage()写入消息conn.WriteMessage()
该服务器只是将所有传入的Websocket消息回显到客户端，因此这说明了如何将Websocket用于全双工通信。

*/
