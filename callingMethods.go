package main

import (
	"fmt"
)

func callingMethods() {
	// call a method
	port := 3000
	_, err := startWebServer(port, 2)

	fmt.Println(err)

}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Webserver starting...")
	// do important things
	fmt.Println("Webserver started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, nil
}
