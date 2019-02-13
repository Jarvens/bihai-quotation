// auth: kunlun
// date: 2019-02-12
// description:
package main

import (
	"common"
	"fmt"
	"reflect"
)

func main() {

	var target = []string{"a", "b", "c", "d"}

	fmt.Println(reflect.ValueOf(&target))
	fmt.Println(reflect.ValueOf(target))
	fmt.Println(reflect.TypeOf(target))
	fmt.Println(reflect.TypeOf(target).Kind())
	fmt.Printf(" \n%s", common.LOGO)
}
