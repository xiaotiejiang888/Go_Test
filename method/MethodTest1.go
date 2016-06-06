package main

import (
	"fmt"
	"math"
)

type VertexInMethod struct {
	X,Y float64
}

func (v *VertexInMethod) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main() {
	v := &VertexInMethod{3,4}
	fmt.Println(v.Abs())
}
