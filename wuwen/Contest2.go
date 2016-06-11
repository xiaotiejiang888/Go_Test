package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("runtime.NumCPU() = ",runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}//同步包
	wg.Add(10)
	//c:=make(chan bool,10)
	for i:=0;i<10;i++{
		go Go(&wg,i)
	}
	//for i:=0;i<10;i++{
	//	<-c
	//}
	wg.Wait()
}
func Go(wg *sync.WaitGroup,index int)  {
	a := 1
	for i:=0;i<100000000;i++{
		a+=i
	}
	fmt.Println(index,a)
	wg.Done()
}


