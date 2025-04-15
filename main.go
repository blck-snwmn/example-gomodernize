package main

import "fmt"

func main() {
	a := 10
	b := 20

	var minVal int
	if a < b {
		minVal = a
	} else {
		minVal = b
	}

	fmt.Println("Minimum:", minVal)
}
