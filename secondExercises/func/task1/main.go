package main

import "fmt"

func name(m map[string]string, key string, value string) {
	m[key] = value
}

func main() {
	m := make(map[string]string)
	name(m, "Alik", "alchemist")
	fmt.Println(m)
}
