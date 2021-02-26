package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999) // create a byte slice with an arbitrary large number of spaces
	resp.Body.Read(bs)        // take the response body and use the Reader .Read interface to feed the bs the Body data
	fmt.Println(string(bs))   // convert bytes to a string and print
}
