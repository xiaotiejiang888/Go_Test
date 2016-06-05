package main

import "fmt"

func main()  {
	i,j:=42,2701
	p := &i//指向i的指针赋值给p
	fmt.Println(*p)//打印指针指向的底层的值
	*p = 21//通过指针p设置i
	fmt.Println(i)
	p = &j//指向j的指针赋值给p
	*p = *p / 37
	fmt.Println(j)
}
