package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// key := strconv.Itoa(i)
		key := "10"
		go sendReq(key, wg)
	}
	wg.Wait()
	fmt.Println("Done!")
}

func sendReq(key string, wg *sync.WaitGroup) string {
	defer wg.Done()
	reqURL := "http://localhost:8080/get_something"
	query := url.Values{
		"key": []string{key},
	}
	resp, err := http.Get(reqURL + "?" + query.Encode())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("resp: %+v\n", string(body))
	return string(body)
}
