package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	fmt.Println(url.QueryEscape(os.Args[1]))
}
