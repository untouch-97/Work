package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type struc struct {
	name string
	age  int
}

func main() {
	maps := make(map[int]int)
	maps[12] = 12
	fmt.Printf(
		"| map  %d, map %v,|\n",
		reflect.ValueOf(maps).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&maps)),
	)
	var slice []string
	fmt.Printf(
		"| slice  %d, slice %v,|\n",
		reflect.ValueOf(slice).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&slice)),
	)
	//var struc struc
	//fmt.Printf(
	//	"| map  %d, map %v,|\n",
	//	reflect.ValueOf(struc).Pointer(),
	//	*(*reflect.SliceHeader)(unsafe.Pointer(&struc)),
	//)

}
