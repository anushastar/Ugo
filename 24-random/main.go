package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("the variable name is x and its value is: %v./n", x)
	if x >= 0 && x <= 100 {
		fmt.Printf("x is between 0 and 100, %v./n", x)
	} else if x >= 101 && x <= 200 {
		fmt.Printf("x is between 101 and 200, %v./n", x)

	} else {
		fmt.Printf("x is between 201 and 250, %v./n", x)
	}

}
