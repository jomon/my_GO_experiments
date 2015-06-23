package main

import (
	"fmt"
)

func main() {
	// x := []int{1, 2, 3, 4, 5}
	y := make([]int, 3, 9)
	fmt.Println(y)
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
	z := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	key := z[0]
	for i := 0; i < len(z); i++ {

		if z[i] < key {
			key = z[i]
		}
	}
	fmt.Println(key)
}
