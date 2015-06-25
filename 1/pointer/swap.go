package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	t := new(int)
	*t = *x
	*x = *y
	*y = *t
}
func main() {
	x, y := 5, 10
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
