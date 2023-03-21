package main

import "fmt"

// цикл, который выведет числа от 3 до 99, но только те, что делятся на 3

func main() {
	for i := 3; i <= 99; i += 3 {
		fmt.Println(i)
	}
}
