package main

import (
	"fmt"
)

// цикл, от 0 до 99, который на четной итерации будет выводить, "ok", а на нечетной "неок"

func main() {
	for i := 0; i < 99; i++ {
		if i%2 == 0 {
			fmt.Println("ok")
		} else {
			fmt.Println("neok")
		}
	}
}
