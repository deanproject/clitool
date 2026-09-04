package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const requestTimeout = 5 * time.Second

func urlhealth() {

	//url := "https://www.fpwop.com/feed/"

	url := []string{"https://linkedin.com", "https://google.com", "https://fb.com", "https://fwppe.com"}

	var wg sync.WaitGroup
	client := http.Client{Timeout: requestTimeout}

	for _, urlStr := range url {

		wg.Add(1)

		go func(urlStr string) {
			defer wg.Done()
			resp, err := client.Get(urlStr)
			if err != nil {
				fmt.Printf("URL: %s | Status: DOWN (%s)\n", urlStr, err.Error())
				return
			}
			defer resp.Body.Close()

			status := "DOWN"
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				status = "UP"
			}
			fmt.Printf("URL: %s | Status: %s (%s)\n", urlStr, status, resp.Status)
		}(urlStr)

	}
	wg.Wait()
}
