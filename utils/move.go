package utils

import (
	"fmt"
	"strings"

	"lemin/models"
)

func MoveAnts(paths []models.Path) []string {
	var line []string
	var movements []string
	scope := len(paths[len(paths)-1].Rooms) + 1 // Length of longest path, add 1 to include start room

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

// Prints out ant movement to terminal
func SimulateMovement(moves []string) {
	for _, move := range moves {
		fmt.Println(move)
	}
}
