package main

/*
...
goroutine:95 is doing job:95
goroutine:96 is doing job:96
goroutine:97 is doing job:97
goroutine:98 is doing job:98
goroutine:99 is doing job:99
all done
duration:  115.380375ms
*/
import (
	"fmt"
	"strconv"
	"time"
)

const (
	numJobs = 100
)

func doJob(name string, i int) {
	time.Sleep(time.Duration(time.Millisecond))
	fmt.Printf("goroutine:%d is doing job:%s \n", i, name)
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("duration: ", time.Since(start))
	}()

	for i := 0; i < numJobs; i++ {
		jobName := strconv.Itoa(i)
		doJob(jobName, i)
	}

	fmt.Println("all done")
}
