package main

import "fmt"
func main() {
	pow := make([]int, 10)
	for i :=range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value :=range pow{//只获取value值
		fmt.Printf("%d====\n",value)
	}
	for value :=range pow{//只获取索引值
		fmt.Printf("%d++++\n",value)
	}
}
