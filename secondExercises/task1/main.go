package main

import "fmt"

func main() {
	var a []int //выделить память для слайса
	for i := 1; i <= 50; i++ {
		a = append(a, i*2)
	}
	fmt.Println(a)
}
