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

func AssignAnts(ants int, paths []models.Path) {
	var minCostIndex int // Set default value for minimum cost

	// Check for no-quantity inputs
	if ants == 0 || len(paths) == 0 {
		return
	}

	// Sort paths by length, shortest first
	SortPaths(paths)

	// Update cost field of individual paths
	UpdateCost(paths)

	// Initialize ant list, assign ids
	antList := make([]Ant, ants)
	for i := range antList {
		antList[i] = Ant{
			Id:       i + 1,
			Position: -1,
		}
	}

	// Assign ants to paths, update their total and cost fields
	for i := 0; i < ants; i++ {
		minCostIndex = FindMinCostInIndex(paths)
		currAnt := fmt.Sprintf("%d", antList[i].Id)

		paths[minCostIndex].Ants = append(paths[minCostIndex].Ants, currAnt)
		paths[minCostIndex].TotalAnts++
		paths[minCostIndex].Cost++
	}
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

func FindMinCostInIndex(paths []models.Path) int {
	var minCost int
	for i := range paths {
		if paths[i].Cost < paths[minCost].Cost {
			minCost = i
		}
	}
	return minCost
}

func UpdateCost(paths []models.Path) {
	for i := range paths {
		cost := len(paths[i].Rooms) + len(paths[i].Ants)
		paths[i].Cost = cost
	}
}
