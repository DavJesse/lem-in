package utils

import (
	"fmt"
	"lemin/models"
	"strings"
)

func MoveAnts(paths []models.Path) []string {
	var line []string
	var movements []string
	scope := len(paths[len(paths)-1].Rooms)

	for pathInd := 1; pathInd <= scope; pathInd++ {

		for _, path := range paths {
			for antInd, ant := range path.Ants {
				position := pathInd - antInd - 1

				if position >= 0 && position < len(path.Rooms) {
					line = append(line, fmt.Sprintf("L%v - %v", ant, path.Rooms[position]))
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
