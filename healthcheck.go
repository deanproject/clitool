package main

import (
	"fmt"
	"net/http"
	"sync"
)

func urlhealth() {

	//url := "https://www.fpwop.com/feed/"

	var wg sync.WaitGroup

	url := [4]string{"https://linkedin.com", "https://google.com", "https://fb.com", "https://fwppe.com"}

	for _, urlresponse := range url {

		wg.Add(1)

		go func(urlStr string) {
			defer wg.Done()
			resp, err := http.Get(urlresponse)
			if err != nil {
				fmt.Println("Error", err.Error())
				return
			}
			defer resp.Body.Close()

			fmt.Printf("URL: %s | Status: %s\n", urlStr, resp.Status)
		}(urlresponse)

	}
	wg.Wait()
}
