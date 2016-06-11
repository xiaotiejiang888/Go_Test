package main

import "fmt"

func main() {
	closure := closure(10)
	fmt.Println(closure(1))
	fmt.Println(closure(2))
}

func closure(x int) func(int) int{
	fmt.Printf("%p\n",&x)
	return func(y int) int {
		fmt.Printf("%p\n",&x)
		return x+y
	}
}
