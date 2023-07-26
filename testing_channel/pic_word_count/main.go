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
	"time"
)

const (
	numJobs = 100
)

func getImgs() []string {
	imgs := []string{}
	for i := 0; i < numJobs; i++ {
		imgs = append(imgs, strconv.Itoa(i))
	}
	return imgs
}

func countWords(img string, i int) (int, error) {
	time.Sleep(time.Duration(time.Millisecond))
	fmt.Printf("goroutine:%d is doing job:%s \n", i, img)
	if i%11 == 0 { // pass some jobs
		return 0, fmt.Errorf("intended error")
	}
	return i, nil
}
func main() {
	start := time.Now()
	defer func() {
		fmt.Println("duration: ", time.Since(start))
	}()
	total := 0
	imgs := getImgs()
	ch := make(chan int, len(imgs))
	// var wg sync.WaitGroup
	for i, imgURL := range imgs {
		i := i
		imgURL := imgURL
		go func() {
			c, err := countWords(imgURL, i)
			if err != nil {
				ch <- 0
				return
			}
			ch <- c
		}()
	}
	fmt.Println(len(ch))
	// for c := range ch {
	for i := 0; i < len(imgs); i++ {
		c := <-ch
		total += c
	}
	fmt.Println("all done ", total)
}
