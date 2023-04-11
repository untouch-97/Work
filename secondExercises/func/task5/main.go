package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func reads(m map[person]bool, key person) bool {
	value, _ := m[key]
	return value
}

func main() {
	m := map[person]bool{
		{"Alik", 25}: true,
	}
	key := person{"Alik", 25}
	fmt.Println(reads(m, key))

}
