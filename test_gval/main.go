package main

import (
	"fmt"

	"github.com/PaesslerAG/gval"
)

func main() {

	value, err := gval.Evaluate("foo > 0", map[string]interface{}{
		"foo": -1.,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(value)

}
