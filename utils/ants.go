package utils

import "sort"

type Ant struct {
	Id       int
	PathIdx  int
	Position int
}

type Paths [][]string

func MoveAnts(ants int, paths Paths) [][]string {
	return paths
}

func SortPaths(paths [][]string) {
	sort.Sort(Paths(paths))
}

func (p Paths) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Paths) Len() int {
	return len(p)
}

func (p Paths) Less(i, j int) bool {
	// Sort by the first element of each inner slice
	// Handle cases where an inner slice might be empty
	if len(p[i]) == 0 {
		return true
	}
	if len(p[j]) == 0 {
		return false
	}
	return p[i][0] < p[j][0]
}
