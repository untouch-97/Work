package main

import "fmt"

// цикл, который пробежится 50 раз, и выведет четные числа каждый раз больше чем предыдущее

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(i * 2)
	}
}
