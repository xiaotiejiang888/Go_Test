package main

import "fmt"

type VertexMap struct {
	Lat,Long float64
}

var m map[string]VertexMap

func main()  {
	m = make(map[string]VertexMap)
	//m = nil
	m["zmy"] = VertexMap{
		40.1111,-70.444,
	}
	fmt.Println(m["zmy"])
}