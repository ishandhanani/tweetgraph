package main

import "fmt"

const (
    twitterBaseURL = "https://api.twitter.com/2/users/%s/bookmarks"
    maxResults     = 100  // Maximum results per page
)

func main() {
	fmt.Println("Hello, World!")
	populate()
}