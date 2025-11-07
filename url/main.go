package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "https://example.com:8080/path/to/resource?key1=value&key2=value2"

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("Can't Parse URL", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)

	parsedURL.Path = "/newpath"
	parsedURL.RawQuery = "username=samixel"

	newURL := parsedURL.String()

	fmt.Println("new URL :", newURL)
}
