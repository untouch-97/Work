package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	defer fmt.Println(v)
	p := &v
	defer fmt.Println(p)
	p.X = 1e9
	defer fmt.Println(v, *p)
}
