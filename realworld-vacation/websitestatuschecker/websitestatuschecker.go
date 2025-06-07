package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	websites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.nonexistentwebsitexyz123.com",
		"https://httpstat.us/503",
	}

	//create channel as string

	c := make(chan string)

	//for loop sending in channel
	for _, site := range websites {
		go checkwebsite(site, c)
	}

	//for loop for receiving in channel

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}

}

func checkwebsite(url string, c chan string) {

	//create client
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	//// handle error (e.g., site is down, DNS failure, etc.)
	resp, err := client.Get(url)
	if err != nil {
		c <- fmt.Sprintf(" %s error is %v", url, err)
		return
	}

	defer resp.Body.Close()
	//// check if status code is 200 (OK) or something else (like 503)
	if resp.StatusCode == 200 {
		c <- fmt.Sprintf("%s website is up", url)
	} else {
		c <- fmt.Sprintf("%s is %d", url, resp.StatusCode)
	}

}
