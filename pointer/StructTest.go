package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{4,6}
	v.X = 9
	fmt.Println(v.X)
}
