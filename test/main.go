package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Printf("33 example-test-repo: %s", time.Now().Format(time.RFC3339))
		time.Sleep(3 * time.Second)
	}
}
