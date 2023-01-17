package main

import (
	"fmt"

	"github.com/Eric-lab-star/hello/morestrings"
)

func main() {
	r := morestrings.ReverseRunes("Hello, world")
	fmt.Println(r)
}
