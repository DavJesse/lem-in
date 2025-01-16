package utils

import (
	"lemin/models"
)

type Ant struct {
	Id       int
	PathIdx  int
	Position int
}

func MoveAnts(ants int, paths []models.Path) [][]string {
	var moves [][]string
	// Sort paths by length, shortest first
	SortPaths(paths)

	// Initialize ant list, assign ids
	antList := make([]Ant, ants)
	for i := range antList {
		antList[i] = Ant{
			Id:       i + 1,
			Position: -1,
		}
	}

	for {
		finished := true
		var currentMoves []string

		// for i := range antList {
		// 	if antList[i].Position < len(paths[0])-1 {
		// 		finished = false
		// 		antList[i].Position++
		// 		move := fmt.Sprintf("L%d-%s", antList[i].Id, paths[0][antList[i].Position])
		// 		currentMoves = append(currentMoves, move)
		// 	}
		// }

		if finished {
			break
		}
		moves = append(moves, currentMoves)
	}

	return moves
}

func SortPaths(paths []models.Path) {
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i].Rooms) > len(paths[j].Rooms) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
}
