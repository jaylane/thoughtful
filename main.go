package main

import (
	"fmt"

	stacks "github.com/jaylane/thoughtful/stacks"
)

func sort(width, height, length, mass int) string {
	volume := width * height * length
	isBulky := volume >= 1000000 || width >= 150 || height >= 150 || length >= 150
	isHeavy := mass >= 20

	if isBulky && isHeavy {
		return stacks.REJECTED
	} else if isBulky || isHeavy {
		return stacks.SPECIAL
	} else {
		return stacks.STANDARD
	}
}

func main() {
	// Test cases
	fmt.Println(sort(100, 100, 100, 10)) // STANDARD
	fmt.Println(sort(200, 100, 100, 10)) // SPECIAL
	fmt.Println(sort(100, 100, 100, 25)) // SPECIAL
	fmt.Println(sort(200, 200, 200, 25)) // REJECTED
}
