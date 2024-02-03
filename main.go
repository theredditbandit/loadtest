package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var url = flag.String("u", "", "url")
var request = flag.Int("n", 10, "number of requests")
var concurrent = flag.Int("c", 0, "number of concurrent requests")

func main() {
	flag.Parse()
	if *url == "" {
		fmt.Println("Need to provide a url with -u flag")
	}
	makeReq(*url, *request, *concurrent)
}

func makeReq(url string, n int, concurrent int) {
	c := false
	wg := sync.WaitGroup{}
	wg.Add(concurrent)
	if concurrent != 0 {
		c = true
	}
	if !c {
		fmt.Println("gothere")
		for i := 0; i < n; i++ {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			defer resp.Body.Close()
			fmt.Printf("Response code: %v\n", resp.Status)
		}
	} else {
		for i := 0; i < n; i++ {
			go func() {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				defer resp.Body.Close()
				fmt.Printf("Response code: %v\n", resp.Status)
			}()
		}
		wg.Wait()
	}
}
