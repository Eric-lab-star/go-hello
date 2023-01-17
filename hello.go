package main

import (
	"fmt"

	"github.com/Eric-lab-star/go-hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	r := morestrings.ReverseRunes("Hello, world")
	fmt.Println(r)
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
