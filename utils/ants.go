package utils

import (
	"fmt"

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

	for i := 0; i < len(paths); i++ {
		precursor := []string{}
		moves = append(moves, precursor)
	}

	var currPath int
	for i := 0; i < ants; i++ {
		// Establish variables
		var nextPath int
		if currPath == len(paths)-1 {
			nextPath = 0
		} else {
			nextPath = currPath + 1
		}

		roomsInCurrPath := len(paths[currPath].Rooms)
		roomsInNextPath := len(paths[nextPath].Rooms)
		antsInCurrPath := paths[currPath].Ants

		// Establish parameter to decide where to move ant
		if (roomsInCurrPath + antsInCurrPath) > roomsInNextPath {
			roomIndex := paths[nextPath].Indicator

			// Send ant to next path
			pathMove := fmt.Sprintf("L%d-%s", antList[i].Id, paths[nextPath].Rooms[roomIndex])
			moves[nextPath] = append(moves[nextPath], pathMove)
			if paths[nextPath].Indicator < roomsInNextPath-1 {
				paths[nextPath].Indicator++
			}
			paths[nextPath].Ants++
			currPath = nextPath

		} else {
			roomIndex := paths[currPath].Indicator

			// Send ant to current path
			pathMove := fmt.Sprintf("L%d-%s", antList[i].Id, paths[currPath].Rooms[roomIndex])
			moves[currPath] = append(moves[currPath], pathMove)

			paths[currPath].Ants++
			if paths[currPath].Indicator < roomsInCurrPath-1 {
				paths[currPath].Indicator++
			}
		}
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
