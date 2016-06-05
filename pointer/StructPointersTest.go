package main

import "fmt"

type  VertexNew struct {
	X int
	Y int
}
var(
	v1 = VertexNew{1,2}//类型为Vertex1
	v2 = VertexNew{X: 1}//Y:0 被省略
	v3 = VertexNew{}//X:0和Y:0
	p=&VertexNew{1,2}//类型为 *Vertex1
)

func main()  {
	//v := Vertex1{1,2}
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)
	fmt.Println(v1,p,v2,v3)
}
