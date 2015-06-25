package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func cir(c *circle) float64 {

	return math.Pi * c.r * c.r
}

func main() {
	c := circle{10.0}
	fmt.Println(cir(&c))
}
