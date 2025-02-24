package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return fmt.Errorf("something went wrong when trying to exit")
}
