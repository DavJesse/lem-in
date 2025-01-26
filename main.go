package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"lemin/utils"
)

func moveAnts(ants int, paths [][]string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	// Sort paths by length
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}

	// Initialize ants
	type ant struct {
		id       int
		pathIdx  int
		position int
	}

	antList := make([]ant, ants)
	for i := range antList {
		antList[i] = ant{
			id:       i + 1,
			pathIdx:  0,
			position: -1,
		}
	}

	var moves [][]string
	for {
		finished := true
		var currentMoves []string

		for i := range antList {
			if antList[i].position < len(paths[0])-1 {
				finished = false
				antList[i].position++
				move := fmt.Sprintf("L%d-%s", antList[i].id, paths[0][antList[i].position])
				currentMoves = append(currentMoves, move)
			}
		}

		if finished {
			break
		}
		moves = append(moves, currentMoves)
	}

	return moves
}

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
