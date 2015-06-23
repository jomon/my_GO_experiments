package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["fe"] = "iron"
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	if name, ok := elements["fe"]; ok {
		fmt.Println(name, ok)
	}
}
