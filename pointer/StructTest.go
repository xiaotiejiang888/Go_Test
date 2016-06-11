package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
type person struct {
	Name string
	Age int
}

func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{4,6}
	v.X = 9
	fmt.Println(v.X)
	a :=person{
		Name:"joe",
		Age:19,
	}
	fmt.Println(a)
	b :=person{"joe111",16}
	fmt.Println(b)
}
