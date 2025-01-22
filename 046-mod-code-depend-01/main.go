package main

import (
	"fmt"

	"github.com/anushastar/puppy"
)

func main() {

	fmt.Println("get the puppy code")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
}
