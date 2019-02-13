// auth: kunlun
// date: 2019-02-13
// description:
package ws

// websocket interface indicate
type Bihai interface {

	// 初始化
	WsInit()

	// 订阅
	Subscribe()

	// 取消订阅
	UnSubScribe()

	// 服务端主动向客户端发起心跳
	Ping()

	// 返回客户端心跳回包
	Pong()
}
