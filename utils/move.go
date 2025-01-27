package utils

import (
	"fmt"
	"strings"

	"lemin/models"
)

func MoveAnts(paths []models.Path, ants int) []string {
	var line []string
	var movements []string
	var scope int
	longestPath := len(paths[len(paths)-1].Rooms) + 1
	maxTurns := (ants * 2) - 2

	// Define scope of trying ant movements
	if longestPath > maxTurns {
		scope = longestPath
	} else {
		scope = maxTurns
	}

	// Try different positions to map valid movements of ants
	for pathInd := 1; pathInd <= scope; pathInd++ {

		for _, path := range paths {
			for antInd, ant := range path.Ants {
				position := pathInd - antInd - 1

				if position >= 0 && position < len(path.Rooms) {
					line = append(line, fmt.Sprintf("L%v-%v", ant, path.Rooms[position]))
				}

			}
		}
		if len(line) > 0 {
			movements = append(movements, strings.Join(line, " "))
			line = nil
		}

	}
	return movements
}

func MostAntsInPath(paths []models.Path) int {
	var max int
	for i, path := range paths {
		if len(path.Ants) > len(paths[max].Ants) {
			max = i
		}
	}

	return len(paths[max].Ants)
}

// Prints out ant movement to terminal
func SimulateMovement(moves []string) {
	for _, move := range moves {
		fmt.Println(move)
	}
}
