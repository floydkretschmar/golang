package website

import (
	"fmt"
	"net/http"
	"time"
)

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func WebsiteChecker() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go CheckLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			CheckLink(link, c)
		}(l)
	}
}
