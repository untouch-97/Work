package main

import (
	"fmt"
)

func main() {
	m := map[string]bool{
		"Alik":    true,
		"Baha":    true,
		"Vitya":   true,
		"Dilshod": true,
		"Dima":    true,
		"Fatima":  true,
		"Ilya":    true,
		"Ivan":    true,
		"Svet":    true,
		"Vasya":   true,
	}

	var name string
	fmt.Println("your name?")
	fmt.Scanln(&name)

	if _, ok := m[name]; ok {
		fmt.Println(name, "is in the list")
	} else {
		fmt.Println(name, "is not in the list")
	}
}
