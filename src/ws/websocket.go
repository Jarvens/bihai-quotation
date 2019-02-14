// auth: kunlun
// date: 2019-02-13
// description:
package ws

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

/**
* websocket
* upgrade 获取connection
* 初始化connection
* 通过chan区分读写通道
* 协程 读 [多路复用 inChan]
* 协程 写 [多路复用 outChan]
*
*
**/
var upgrade = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}, EnableCompression: true, HandshakeTimeout: 10 * time.Second}

type Bihai interface {
	//初始化
	WsInit()

	//订阅
	Subscribe()

	//取消订阅
	UnSubScribe()

	//服务端主动向客户端发起心跳
	Ping()

	//服务端回包客户端
	Pong()

	//读取客户端发送的消息
	ReadMessage()
}

type Connection struct {
	wsCon     *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

func WsHandle(w http.ResponseWriter, r *http.Request) {
	var (
		wsCon *websocket.Conn
		err   error
		conn  *Connection
		data  []byte
	)

	if wsCon, err = upgrade.Upgrade(w, r, nil); err != nil {
		fmt.Printf("wsHandle upgrade error: %v\n", err)
		return
	}

	if conn, err = WsInit(wsCon); err != nil {
		fmt.Printf("wsHandle init error: %v\n", err)
		conn.Close()
	}

	for {
		if _, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}

func WsInit(wsCon *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{wsCon: wsCon,
		inChan:    make(chan []byte, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan byte, 1)}
	conn.onConnected()
	go conn.loopReadMessage()
	go conn.loopWriteMessage()
	return
}

// 读取消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
		address := conn.wsCon.RemoteAddr().String()
		fmt.Printf("打印ReadMessage: %v\n", string(data) == "")
		DispatchMessage(data, address)
		fmt.Printf("receive client message: \n%s\n", string(data))

		// TODO 消息处理器区分 心跳/订阅/取消订阅

	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 写入消息
func (conn *Connection) WriteMessage(data []byte) (err error) {
	return nil
}

func (conn *Connection) Close() {
	log.Printf("客户端[%s]下线了!\n", conn.wsCon.RemoteAddr().String())
	conn.wsCon.Close()
	//开启同步锁，防止并发
	conn.mutex.Lock()
	// TODO 删除客户端连接信息
	if !conn.isClosed {
		//关闭通道，此时chan会接收到 0值，需要特殊处理
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) loopReadMessage() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsCon.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			log.Println("收到closeChan的消息")
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) loopWriteMessage() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) onConnected() {
	fmt.Printf("\n 客户端【%s】加入会话\n", conn.wsCon.RemoteAddr().String())
}
