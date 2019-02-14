// auth: kunlun
// date: 2019-02-13
// description:
package common

const (
	SubSuccess = iota
	SubFailure
)
const (
	UnSubSuccess = iota
	UnSubFailure
	Sub   = "sub"
	UnSub = "unsub"
	Ping  = "ping"
)

type SubRequest struct {
	Id   string `json:"id"`
	Sub  string `json:"sub"`
	Rate int    `json:"rate"`
}

type UnSubRequest struct {
	Id    string `json:"id"`
	UnSub string `json:"unsub"`
}

func NewSubRequest() *SubRequest {
	return &SubRequest{"client1", "market.btcusdt.kline.1m", 0}
}

func NewUnsubRequest() *UnSubRequest {
	return &UnSubRequest{"client2", "market.btcusdt.kline.1m"}
}
