package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func add(m map[person]bool, name string, age int, value bool) {
	m[person{name, age}] = value
}
func main() {
	m := map[person]bool{}
	add(m, "Alik", 25, true)
	fmt.Println(m)
}
