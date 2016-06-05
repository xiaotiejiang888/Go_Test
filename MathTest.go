package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite Num is",rand.Intn(10))//这个程序的运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字。 （为了得到不同的随机数，需要提供一个随机数种子，参阅 rand.Seed。）
	//fmt.Println("随机",rand.Seed(15))
}