package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return fmt.Errorf("Something went wrong when trying to exit")
}
