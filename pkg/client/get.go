package client

import (
	"fmt"
	"log"
	"net/http"
)

func Get(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error making new request:", err)
		log.Fatal(err)
	}

	return req
}
