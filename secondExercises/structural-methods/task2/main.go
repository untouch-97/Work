package main

type Person struct {
	mane    string
	surname string
	age     int
}

func (p *Person) NewPerson(name, surname string, age int) *Person {
	return &Person{
		name,
		surname,
		age,
	}

}
