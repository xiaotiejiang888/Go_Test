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
type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func main() {
	//m :=Manager{User:User{1,"ok",12},title:"123"}
	u := User{1, "ok", 18}
	Set(&u)
	fmt.Println(u)
	//Info(m)//值拷贝
	//Info(&u)//指针传递
	//t := reflect.TypeOf(m)
	////fmt.Printf("%#v\n",t.Field(1))
	//fmt.Printf("%#v\n",t.FieldByIndex([]int{0,1}))
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)//反射出类型名称
	fmt.Println("Type:", t.Name())
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("xx")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)//反射出字段
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)//反射出方法
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr&&!v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String{
		f.SetString("beybey")
	}
}