// auth: kunlun
// date: 2019-02-13
// description:
package ws

import (
	"common"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"utils"
)

var GlobalSubMap SubMap

func DispatchMessage(message []byte, addr string) {
	var target interface{}
	//转换为map
	json.Unmarshal(message, &target)
	refV := reflect.ValueOf(target)
	if refV.MapIndex(reflect.ValueOf(common.Sub)).IsValid() {
		//TODO 订阅
		SubscribeHandle(addr, fmt.Sprintf("%s", refV.MapIndex(reflect.ValueOf(common.Sub))), 0)

	} else if refV.MapIndex(reflect.ValueOf(common.UnSub)).IsValid() {
		//TODO 取消订阅

	} else if refV.MapIndex(reflect.ValueOf(common.Ping)).IsValid() {
		//TODO 客户端心跳

	}
}

// 订阅处理器
// addr    客户端地址
// message 订阅信息 market.btcusdt.kline.1m
// rate    推送频率
func SubscribeHandle(addr string, message string, rate int) {
	values := strings.Split(message, ".")
	symbol := values[1]
	period := values[3]
	contains, _ := utils.Contain(addr, GlobalSubMap.Map)
	if contains {
		option := GlobalSubMap.Map[addr]
		containSymbol, _ := utils.Contain(symbol, option)
		if containSymbol {
			fmt.Println(period)
		} else {

		}
	}

}

// 取消订阅处理器
func UnSubscribeHandle(addr string, message string) {

}

// 心跳处理器
func PingHandle() {

}

type SubOption struct {
	Period string
	Rate   int
}

type SubMap struct {
	Map map[string]map[string]SubOption
}
