package main

import (
	"log"
	"mnp-tests-server/tests"
)

func main() {
	log.Println("Running user tests...")

	if err := tests.RunUserCRUD(); err != nil {
		log.Fatalf("User tests failed: %v", err)
	}

	log.Println("User tests completed successfully!")
}
