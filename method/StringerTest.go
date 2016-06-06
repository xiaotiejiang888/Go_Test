package main

import "fmt"

type Person struct {
	Name string
	Age int
}
func (p Person) String() string{//相当于java的重写toString方法
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}
func main() {
	a := Person{"Albert",26}
	z := Person{"Jerry",32}
	fmt.Println(a,z)
	//{Albert 26} {Jerry 32}
	//Albert (26 years) Jerry (32 years)
}
