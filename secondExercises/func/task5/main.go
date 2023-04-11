package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func reads(m map[person]bool, key person) {
	value, _ := m[key]

}

func main() {
	m := map[person]bool{
		{"Alik", 25}: true,
	}
	key := person{"Alik", 25}

}
