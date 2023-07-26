package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	numJobs       = 100
	numGoroutines = 2
)

func doJob(ch chan string, i int) {
	for name := range ch {
		time.Sleep(time.Duration(time.Millisecond))
		fmt.Printf("goroutine:%d is doing job:%s \n", i, name)
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("duration: ", time.Since(start))
	}()

	signalCh := make(chan struct{}, numJobs)
	workCh := make(chan string, numJobs)

	for i := 0; i < numJobs; i++ {
		jobName := strconv.Itoa(i)
		go func() {
			workCh <- jobName
			signalCh <- struct{}{}
		}()
	}

	for j := 0; j < numGoroutines; j++ {
		j := j
		go func() {
			doJob(workCh, j)
			fmt.Println("job done: ", j)
		}()
	}
	for i := 0; i < numJobs; i++ {
		<-signalCh
	}
	fmt.Println("all done")
}
