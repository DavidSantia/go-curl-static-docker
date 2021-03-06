package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, "http://www.example.com/")
		easy.Perform()
	}
	fmt.Printf("\n")
}
