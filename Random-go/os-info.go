package main

import (
	"fmt"
	"os"
)

func getHostName() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n", hostname)
}

func main() {
	getHostName()
}
