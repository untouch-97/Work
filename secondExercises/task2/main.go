package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]string)
	m["name1"] = "Alik"
	m["name2"] = "Baha"
	m["name3"] = "Vitya"
	m["name4"] = "Dilshod"
	m["name5"] = "Dima"
	m["name6"] = "Fatima"
	m["name7"] = "Ilya"
	m["name8"] = "Ivan"
	m["name9"] = "Svet"
	m["name10"] = "Vasya"

	var n string
	fmt.Println("your name?")
	fmt.Scan(&n)

	if m[strings.ToLower(n)] != "" {
		fmt.Println("Correct")
	} else {
		fmt.Println("Wrong")
	}
}
