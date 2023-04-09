package main

import (
	"fmt"
)

func name(m map[string]string, key string, value string) {
	m[key] = value
}

func fn(m map[string]string, key string) string {
	value, _ := m[key]
	return value
}

func yesORno(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

func main() {
	m := make(map[string]string)
	staff := make(map[string]person)
	name(m, "Alik", "manager")
	fmt.Println(fn(m, "Alik"))
	fmt.Println(staff)
}

type person struct {
	name string
	age  int
}

func add(staff map[string]person, jobTitle string, value person) {
	staff[jobTitle] = value
}

func valueStaff(staff map[string]person, key string) person {
	return staff[key]
}
func boolStaff(staff map[string]person, jobTitle string) bool {
	_, ok := staff[jobTitle]
	return ok
}
