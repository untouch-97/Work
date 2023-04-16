package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetSurName(surname string) {
	p.surname = surname
}

func (p *Person) SetAge(age int) {
	p.age = age
}
func main() {
	list := Person{"Bahodir", "Khayrullaev", 25}
	list.SetName("Suleyman")
	fmt.Println(list)
}
