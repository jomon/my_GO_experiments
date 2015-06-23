package main

import (
	"fmt"
)

func fact(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * fact(x-1)
	}

}
func main() {
	fmt.Println("enter the number:")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println(fact(input))

}
