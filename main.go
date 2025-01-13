package main

import (
	"fmt"
	"os"

	"lemin/input"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}

	// Parse input
	ants, rooms, links, err := input.ParseInput(os.Args[1])
	fmt.Println(ants)
	fmt.Println(rooms)
	fmt.Println(links)
	fmt.Println(err)
}
