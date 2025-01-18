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

func AssignAnts(ants int, paths []models.Path) []models.Path {
	// Check for no-quantity inputs
	if ants == 0 || len(paths) == 0 {
		return paths
	}
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
		antsInCurrPath := paths[currPath].TotalAnts

		// Establish parameter to decide where to move ant
		if (roomsInCurrPath + antsInCurrPath) > roomsInNextPath {

			// Send ant to next path
			currAnt := fmt.Sprintf("%d", antList[i].Id)
			paths[nextPath].Ants = append(paths[nextPath].Ants, currAnt)

			paths[nextPath].TotalAnts++
			currPath = nextPath

		} else {

			// Send ant to current path
			currAnt := fmt.Sprintf("%d", antList[i].Id)
			paths[currPath].Ants = append(paths[currPath].Ants, currAnt)

			paths[currPath].TotalAnts++
		}
	}

	return paths
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
