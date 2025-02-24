package main

import "fmt"

func commandMap() error {
	locations := []string{"place 1", "place 2"}
	for _, location := range locations {
		fmt.Printf("%s\n", location)
	}

	return nil
}
