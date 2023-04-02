package main

import "fmt"

//	type Vertex struct {
//		X int
//		Y int
//	}
//
//	func main() {
//		v := Vertex{1, 2}
//		defer fmt.Println(v)
//		p := &v
//		defer fmt.Println(p)
//		p.X = 1e9
//		defer fmt.Println(v, *p)
//	}
type list struct {
	name1, name2, name3, name4, name5, name6, name7, name8, name9, name10 string
}

func main() {
	m := make(map[string]list)
	m["names"] = list{name1: "Alik", name2: "Baha", name3: "Vitya", name4: "Dilshod", name5: "Dima", name6: "Fatima", name7: "Ilya", name8: "Ivan", name9: "Svet", name10: "Vasya"}

	var n string
	fmt.Println("your name?")
	fmt.Scanln(&n)

	if _, ok := m[n]; ok {
		fmt.Println(n, "is in the list")
	} else {
		fmt.Println(n, "is not in the list")
	}
}
