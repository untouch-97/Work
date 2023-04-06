package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m["names"] = []string{"Alik", "Baha", "Vitya", "Dilshod", "Dima", "Fatima", "Ilya", "Ivan", "Sveta", "Vasya"}

	fmt.Println("your name?")
	var name string
	fmt.Scanln(&name)

	for _, value := range m["names"] {
		switch {
		case value == name:
			fmt.Println(name, "is in the list")
			return
		}
	}
	fmt.Println(name, "is not in the list")
}
