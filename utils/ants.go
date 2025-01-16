package utils

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
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
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
