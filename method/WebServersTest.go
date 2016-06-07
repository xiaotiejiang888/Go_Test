package main

import (
	"net/http"
	"fmt"
	"log"
)

type Hello struct {}

func (h Hello) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w,"Hello!")
}
func main() {
	var h Hello
	error := http.ListenAndServe("localhost:4000", h)
	if error != nil {
		log.Fatal(error)
	}
}