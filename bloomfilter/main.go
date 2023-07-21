package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	filter.Add([]byte("Love"))
	fmt.Println(filter.Test([]byte("Love")))
	fmt.Println(filter.Test([]byte("Love1")))
}
