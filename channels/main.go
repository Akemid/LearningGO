package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://sergiomondragon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link,c)
	}
	// Block routine until channel receives a message
	// Blocking call <-c

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l:= range c {
		go func(link string){
			time.Sleep(5*time.Second)
			checkLink(link,c)		
		}(l)
	}
	
}

func checkLink(link string, c chan string) {
	// Make a http request
	_, err := http.Get(link)
	// Check if the site is up or down
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	// Send a status report
	fmt.Println(link, "is up!")
	c <- link
}