package main

/*
test 10 goroutines doing 1000 jobs
*/
import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

const (
	numGoroutines = 2
	numJobs       = 1000
)

func doJob(name string, i int) {
	fmt.Printf("goroutine:%d  is doing job:%s \n", i, name)
}

func main() {

	ch := make(chan string)
	for i := 0; i <= numJobs; i++ {
		str := strconv.Itoa(i)
		go func() {
			ch <- str
		}()
	}

	for i := 0; i < numGoroutines; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for name := range ch {
				doJob(name, i)
			}
			fmt.Printf("goroutine:%d  is done \n", i)
		}()
	}
	wg.Wait()
	fmt.Println("all done")
}
