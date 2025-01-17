package utils

import (
	"fmt"
	"log"

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

	// Move the first ant to the first path
	paths[0].Ants++
	curr := []string{}
	move := fmt.Sprintf("L%d-%s", antList[0].Id, paths[0].Rooms[0])
	curr = append(curr, move)
	moves = append(moves, curr)
	log.Printf("Nummber of paths: %d", len(paths))

	for i := 1; i < len(paths); i++ {
		if i < len(paths)-1 {
			pathRooms := len(paths[i].Rooms)
			nextPathRooms := len(paths[i+1].Rooms)
			pathAnts := paths[i].Ants
			param := pathRooms + pathAnts

			if param > nextPathRooms {
				paths[i+1].Ants++
				curr = nil
				move = fmt.Sprintf("L%d-%s", antList[i].Id, paths[i+1].Rooms[i-1])
				log.Printf("move: %s", move)
				curr = append(curr, move)
				moves = append(moves, curr)
			} else {
				paths[i].Ants++
				curr = nil
				move = fmt.Sprintf("L%d-%s", antList[i].Id, paths[i].Rooms[i])
				log.Printf("move: %s", move)
				curr = append(curr, move)
				moves = append(moves, curr)
			}
		}
	}

	// for {
	// 	finished := true
	// 	var currentMoves []string

	// 	for i := range antList {
	// 		if antList[i].Position < len(paths[0])-1 {
	// 			finished = false
	// 			antList[i].Position++
	// 			move := fmt.Sprintf("L%d-%s", antList[i].Id, paths[0][antList[i].Position])
	// 			currentMoves = append(currentMoves, move)
	// 		}
	// 	}

	// 	if finished {
	// 		break
	// 	}
	// 	moves = append(moves, currentMoves)
	// }

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
