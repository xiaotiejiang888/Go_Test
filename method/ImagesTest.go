package main

import (
	"image"
	"fmt"
)

func main() {
	rgba := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(rgba.Bounds())
	fmt.Println(rgba.At(0,0).RGBA())
}
