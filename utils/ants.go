package utils

func MoveAnts(ants int, paths [][]string) [][]string {
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
