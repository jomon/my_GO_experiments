package main

import (
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("joe")
}
