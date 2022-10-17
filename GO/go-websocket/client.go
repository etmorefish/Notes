package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func main() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	socketUrl := "ws://localhost:8888" + "/socket"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)

	// Our main loop for the client
	// We send our relevant packets here
	for {
		select {
		case <-time.After(time.Duration(1) * time.Millisecond * 1000):
			// Send an echo packet every second
			err := conn.WriteMessage(websocket.BinaryMessage, []byte("Hello from GolangDocs!"))
			if err != nil {
				log.Println("Error during writing to websocket:", err)
				return
			}

		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}

/*
这个简单的客户端将每隔1秒钟不断发出消息。如果我们的整个系统按预期工作，
则服务器将接收间隔为1秒的数据包，并回复相同的消息。
客户端还将具有接收传入的Websocket数据包的功能。在我们的程序中，我们将
有一个单独的goroutine处理程序receiveHandler，用于侦听这些传入的数据包。

如果您观察代码，您会发现我创建了两个通道done，interrupt用于receiveHandler()和之间的通信main()。
我们使用无限循环使用select来通过通道监听事件。我们conn.WriteMessage()每秒钟写一条消息。
如果激活了中断信号，则所有未决的连接都将关闭，并且我们可以正常退出！
嵌套select是为了确保两件事：
• 如果receiveHandler通道退出，则通道'done'将关闭。这是第一个case <-done条件
• 如果'done'通道未关闭，则在1秒钟后会有超时，因此程序将在1秒钟超时后退出
通过使用通道仔细处理所有情况select，您可以拥有一个可以轻松扩展的最小体系结构。

*/
