package main

import (
	"fmt"
	"net/http"
)

func urlhealth() {

	//url := "https://www.fpwop.com/feed/"

	url := [4]string{"https://linkedin.com", "https://google.com", "https://fb.com", "https://fwppe.com"}

	for _, urlresponse := range url {

		resp, err := http.Get(urlresponse)
		if err != nil {
			fmt.Println("Error", err.Error())
			continue
		}

		fmt.Printf("URL: %s | Status: %s\n", urlresponse, resp.Status)
		resp.Body.Close()
	}

}
