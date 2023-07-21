package main

/*
test 10 goroutines doing 1000 jobs
*/
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const numGoroutines = 2

func doJob(name string, i int) {
	//defer wg.Done()
	fmt.Printf("goroutine:%d  is doing job:%s \n", i, name)
}

func main() {

	ch := make(chan string)
	for i := 0; i <= 1000; i++ {
		//i := i
		//wg.Add(1)
		//go func() {
		ch <- i[0]
		//}()
	}

	// channel blocked, the following code will never be reached

	for i := 0; i < numGoroutines; i++ {
		i := i
		go func() {
			for name := range ch {
				doJob(name, i)
			}
			fmt.Printf("goroutine:%d  is done \n", i)
		}()
	}
	//wg.Wait()
	fmt.Println("all done")
}
