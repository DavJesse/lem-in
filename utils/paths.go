package utils

import "lemin/models"

func FindPaths(startRoom string, endRoom string, links []models.Link) []models.Path {
	// Initiate utility variables
	var paths []models.Path
	visited := make(map[string]bool)
	currentPath := []string{startRoom}

	var dfs func(current string)
	dfs = func(current string) {
		if current == endRoom {
			var newPath models.Path
			precursor := make([]string, len(currentPath))
			copy(precursor, currentPath)
			newPath.Rooms = precursor
			paths = append(paths, newPath)
			return
		}

		visited[current] = true
		for _, link := range links {
			next := ""
			if link.From == current && !visited[link.To] {
				next = link.To
			} else if link.To == current && !visited[link.From] {
				next = link.From
			}
			if next != "" {
				currentPath = append(currentPath, next)
				dfs(next)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited[current] = false
	}

	dfs(startRoom)
	return paths
}
