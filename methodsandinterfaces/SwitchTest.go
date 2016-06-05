package main

import (
	"fmt"
	"runtime"
	"log"
)

func main() {
	log.Println(runtime.GOOS)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS;os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.",os)
	}
}