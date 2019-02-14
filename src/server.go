// auth: kunlun
// date: 2019-02-12
// description:
package main

import (
	"common"
	"fmt"
	"net/http"
	"ws"
)

func main() {
	inChan := make(chan bool)
	fmt.Printf(" \n%s\n", common.LOGO)
	fmt.Printf("Exchange quotation server starting")
	http.HandleFunc("/", ws.WsHandle)
	go http.ListenAndServe("0.0.0.0:1234", nil)
	fmt.Println(<-inChan)
}
