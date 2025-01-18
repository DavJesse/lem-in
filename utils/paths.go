package utils

import (
	"lemin/models"
)

func FindPaths(startRoom string, endRoom string, links []models.Link) []models.Path {
	var paths []models.Path
	var visited []string
	var current string

	for _, link := range links {
		var path models.Path
		if startRoom == link.From || startRoom == link.To {
			if startRoom == link.From {
				if Discovered(visited, link.From) {
					continue
				}
				visited = append(visited, link.To)
				path.Rooms = append(path.Rooms, link.To)
				current = link.To
			} else {
				if Discovered(visited, link.From) {
					continue
				}
				visited = append(visited, link.From)
				path.Rooms = append(path.Rooms, link.From)
				current = link.From
			}
		}

		if current == link.From || current == link.To {
			if current == link.From {
				if Discovered(visited, link.From) {
					continue
				}
				visited = append(visited, link.To)
				path.Rooms = append(path.Rooms, link.To)
				current = link.To
			} else {
				if Discovered(visited, link.From) {
					continue
				}
				visited = append(visited, link.From)
				path.Rooms = append(path.Rooms, link.From)
				current = link.From
			}
		}

		if endRoom == link.From || endRoom == link.To {
			continue
		}

	}
	return paths
}

func Discovered(visited []string, room string) bool {
	for i := range visited {
		if visited[i] == room {
			return true
		}
	}
	return false
}
