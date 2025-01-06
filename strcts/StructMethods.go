package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")
	rect := Rectangle{
		Height: 100,
		Width:  100,
	}
	area := fmt.Sprintf("area: %v", rect.calculateArea())
	fmt.Println(area)
}

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) calculateArea() int {
	return r.Height * r.Width
}
