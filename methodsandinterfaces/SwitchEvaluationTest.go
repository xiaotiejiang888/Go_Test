package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	log.Println("日志",today)
	switch time.Saturday {
	case today+0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow.")
	case today+2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
