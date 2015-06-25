package main

import (
	"fmt"
	"math"
)

func rect(x, y float64) float64 {
	a := x * y
	return a
}

type circle struct {
	x float64
	y float64
	r float64
}

// c.x = 0.0
// c.y = 0.0
// c.r = 6.0

func circ(c circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := circle{x: 0, y: 0, r: 5}
	fmt.Println(rect(5.0, 6.0))
	fmt.Println(circ(c))
}
