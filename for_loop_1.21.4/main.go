package main

// export GOEXPERIMENT=loopvar; go run main.go

import (
	"fmt"
	"sync"
)

func main() {
	arr := make([]int, 10) // Create an integer array of length 100
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	var wg sync.WaitGroup
	for _, i := range arr {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
