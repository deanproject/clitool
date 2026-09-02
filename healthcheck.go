package main

import (
	"fmt"
	"net/http"
)

func urlhealth() {

	url := "https://www.fpwop.com/feed/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	defer resp.Body.Close()
	fmt.Printf("Status: %d %s\n", resp.StatusCode, resp.Status)

}
