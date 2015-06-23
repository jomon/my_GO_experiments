package main 

import (
	"fmt"
)
func main() {
	
	fmt.Println("entetr number")
	var x int
	fmt.Scanf("%d",&x)
	switch x {
		case 0: fmt.Println("zero")
		case 1: fmt.Println("one")
		case 2: fmt.Println("two")
		case 3: fmt.Println("three")
		case 4: fmt.Println("four")
		case 5: fmt.Println("five")
		default: fmt.Println ("greater than five")
	}
}
