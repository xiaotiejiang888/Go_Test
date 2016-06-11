package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)//默认双向通道
	go func (){
		fmt.Println("Go Go Go")
		//c<-true
		<-c
		//close(c)
	}()
	//<-c//阻塞
	c<-true
	//for v:=range c {
	//	fmt.Println(v)
	//}
	//c <- true//无缓存的时候是有阻塞的（同步阻塞）  有缓存的时候（异步）则不阻塞
	//chanle读和写都有阻塞
}


