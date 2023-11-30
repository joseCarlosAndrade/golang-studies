package main

import (
	"fmt"
	"net/http"
	"bufio" // for reading response content
)

func main() {
	resp,err := http.Get("http://localhost:8080")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Printf("Status: %v \n", resp.StatusCode)

	// reading content response
	scanner := bufio.NewScanner(resp.Body)
	
	for i := 0; scanner.Scan() && i < 5; i++ { // reading the first 5 lines of response or untill theres no more content
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}