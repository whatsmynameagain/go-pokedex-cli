package main

import (
	"fmt"
	"sort"
)

func commandHelp(conf *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	// to avoid displaying the commands in random order (because it's reading from a map)
	// create an array of [commandName, commandDecription] arrays
	var cmdList [][2]string
	// add each command and its description
	for k, v := range getCommands() {
		cmdList = append(cmdList, [2]string{k, v.description})
	}
	// sort the inner arrays by the first element
	sort.Slice(cmdList, func(i, j int) bool {
		return cmdList[i][0] < cmdList[j][0]
	})
	// print commandName: commandDescription
	for _, cmd := range cmdList {
		fmt.Printf("%s: %s\n", cmd[0], cmd[1])
	}
	return nil
}
