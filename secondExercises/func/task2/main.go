package main

import "fmt"

func reads(m map[string]string, key string) string {
	value, _ := m[key]
	return value
}

func main() {
	m := make(map[string]string)
	m["Alik"] = "alchemist"
	fmt.Println(reads(m, "Alik"))
}
