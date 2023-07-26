package main

/*
...

*/
import (
	"fmt"
	"strconv"
	"time"
)

const (
	numJobs = 1000
)

func doJob(ch chan struct{}, name string, i int) {
	time.Sleep(time.Duration(time.Millisecond))
	fmt.Printf("goroutine:%d is doing job:%s \n", i, name)
	ch <- struct{}{}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("duration: ", time.Since(start))
	}()

	ch := make(chan struct{}, numJobs)
	for i := 0; i < numJobs; i++ {
		i := i
		jobName := strconv.Itoa(i)
		// ch <- struct{}{} // wrong, in the outside, ch does not relate to goroutine result
		go doJob(ch, jobName, i)
	}
	for i := 0; i < numJobs; i++ {
		<-ch
	}
	fmt.Println("all done")
}
