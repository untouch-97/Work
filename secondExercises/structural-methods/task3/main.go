package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetSurName() string {
	return p.surname
}
func (p Person) GetAge() int {
	return p.age
}
func main() {
	var list Person
	list = Person{"Bahodir", "Khayrullaev", 25}
	fmt.Println(list.GetName())
	fmt.Println(list.GetSurName())
	fmt.Println(list.GetAge())
}
