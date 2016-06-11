package main

import "fmt"

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)//3 2 1 0
 		defer func() {
			fmt.Println("defer_closure i = ",i) // 不解为何输出4
		}()
		fs[i] = func() {
			fmt.Println("closure i = ",i)//不解为何输出4
		}
	}
	for _,f := range fs{
		f()
	}
	//for i:=0;i<4;i++ {
	//	fmt.Println("i = ",i)
	//}
}
