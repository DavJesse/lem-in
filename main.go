package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"lemin/utils"
)

func main() {
	// Check for valid number command-line arguments
	utils.CheckArgs()

	// Parse input file
	ants, rooms, links, err := utils.ParseInput(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Find start and end rooms
	startRoom, endRoom := utils.FindStartAndEndRooms(rooms)

	// Extract nodes of graph
	nodes := utils.AsignNodes(links)

	// Find all possible paths
	paths := utils.FindPaths(startRoom, endRoom, nodes)
	if len(paths) == 0 {
		log.Fatal("ERROR: no valid path found between start and end")
	}

	// Sort paths by length
	utils.SortPaths(paths)

	// Assign ants to paths
	utils.AssignAnts(ants, paths)
	log.Printf("Paths: %#v", paths)

	// Extract and Print the input data
	content, err := utils.ValidContent(os.Args[1])
	if err != nil {
		log.Fatal("ERROR: invalid data format")
	}

	// Print number of ants and room configuration
	fmt.Println(content[0])
	for _, line := range content[1:] {
		if !strings.Contains(line, "-") {
			fmt.Println(line)
		}
	}

	// Print links
	for _, line := range content[1:] {
		if strings.Contains(line, "-") {
			fmt.Println(line)
		}
	}
	fmt.Println()

	// Extract moves
	moves := utils.MoveAnts(paths)

	// Simuluate ant movement
	utils.SimulateMovement(moves)
}
