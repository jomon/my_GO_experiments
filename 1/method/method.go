package main

import (
	"fmt"
)

type rect struct {
	x, y int
}

func (r *rect) area() int {

	return r.x * r.y
}

func main() {
	r := rect{10, 5}
	fmt.Println(r.area())
}
