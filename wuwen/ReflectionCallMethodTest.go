package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
func (u User) Hello1(name string){
	fmt.Println("Hello",name,", my name is",u.Name)
}

func main() {
	u := User{1, "ok", 18}
	//u.Hello1("zmy")
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello1")
	args :=[]reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
