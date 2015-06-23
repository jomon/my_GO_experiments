package main

import (
	"fmt"
)

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	var x int

	fmt.Println("enter the count:")
	fmt.Scanf("%d", &x)
	z := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Printf("enter the %d number:", i+1)
		fmt.Scanf("%d", &z[i])
	}

	fmt.Println(add(z...))
}
