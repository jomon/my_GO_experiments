package main

import (
	"fmt"
)

func half(z int) (int, bool) {
	x := z / 2
	if ok := z%2 == 0; ok {
		return x, ok
	} else {
		return x, ok
	}
}

func main() {
	fmt.Println("enter the number:")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println(half(input))
}
