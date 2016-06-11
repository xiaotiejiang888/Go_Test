package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main() {
	c, err := redis.Dial("tcp", "192.168.1.103:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
}