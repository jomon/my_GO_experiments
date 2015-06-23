package main

import (
	"fmt"
)

func zero(x *int) {
	*x = 1
}

func main() {
	x := new(int)
	*x = 10
	zero(x)
	fmt.Println(*x)
}
