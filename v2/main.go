package main

import (
	"log"

	"cobra-1176/v2/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
