package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"golang.org/x/sync/singleflight"
)

func main() {
	group := &singleflight.Group{}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := rand.Intn(10)
		key := strconv.Itoa(i)

		fn := func() (interface{}, error) {
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
			return string(body), nil
		}

		go func() {
			key := key
			result, err, shared := group.Do(key, fn)
			fmt.Printf("%+v, %+v, %+v \n", result, err, shared)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done!")
}
