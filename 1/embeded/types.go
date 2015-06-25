package main

import (
	"fmt"
)

type version struct {
	name string
}

type android struct {
	version
	name string
}

func (v *version) print() {

	fmt.Println(v.name)
}
func (a *android) p1() {
	fmt.Println(a.name)

}

func main() {

	a := new(android)
	a.version.name = "123"

	a.version.print()
	a.name = "asd"
	a.p1()
}
