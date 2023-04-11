package main

import "fmt"

func toReturn(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

func main() {
	m := map[string]string{
		"Alik": "alchemist",
	}
	fmt.Println(toReturn(m, "Alik"))
}
