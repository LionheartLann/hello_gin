package main

/*
...
goroutine:644 is doing job:644
goroutine:758 is doing job:758
goroutine:121 is doing job:121
goroutine:938 is doing job:938
goroutine:730 is doing job:730
goroutine:715 is doing job:715
all done
duration:  5.671083ms
*/
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	numJobs = 1000
)

// func doJob(wg sync.WaitGroup, name string, i int) {
// 	defer wg.Done()
// 	time.Sleep(time.Duration(time.Millisecond))
// 	fmt.Printf("goroutine:%d is doing job:%s \n", i, name)
// }

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("duration: ", time.Since(start))
	}()

	var wg sync.WaitGroup
	for i := 0; i < numJobs; i++ {
		i := i
		jobName := strconv.Itoa(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(time.Millisecond))
			fmt.Printf("goroutine:%d is doing job:%s \n", i, jobName)
		}()
	}
	wg.Wait()
	fmt.Println("all done")
}
