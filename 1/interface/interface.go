package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
func main() {
	c, r := new(float64)
	m := new(MultiShape)
	fmt.Println(m.area)
}
