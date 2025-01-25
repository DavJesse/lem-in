package utils

import (
	"log"

	"lemin/models"
)

func FindPaths(startRoom string, endRoom string, nodes map[string][]string) []models.Path {
	// Initialize utility variables
	var paths []models.Path

	// Handle errors related to start and end rooms
	if startRoom == endRoom {
		log.Fatal("ERROR: 'Start' and 'End' rooms identical")
	}

	if _, exists := nodes[startRoom]; !exists {
		log.Fatal("ERROR: 'Start' room not included in source file")
	}

	if _, exists := nodes[endRoom]; !exists {
		log.Fatal("ERROR: 'End' room not included in source file")
	}

	rooms := nodes[startRoom]
	visited := []string{startRoom}

	// Explore rooms liked to start room
	for _, room := range rooms {
		var path models.Path

		// Abandon paths with visited rooms
		if Discovered(visited, room) {
			continue
		}
		
		// Append first room linked to start
		path.Rooms = append(path.Rooms, room)
		if !Discovered(visited, room) {
			visited = append(visited, room)
		}

		// using depth-first search, update path
		path, end := UpdatePath(room, endRoom, &visited, nodes, path)

		if end && path.Rooms != nil {
			paths = append(paths, path)
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

func AsignNodes(links []models.Link) map[string][]string {
	nodes := make(map[string][]string)

	for _, link := range links {
		// Create nodes and links for origin
		rooms1, exists1 := nodes[link.From]
		if exists1 {
			// Avoid duplicating rooms
			if !Discovered(rooms1, link.To) {
				nodes[link.From] = append(nodes[link.From], link.To)
			}
		} else {
			nodes[link.From] = []string{link.To}
		}

		// Create nodes and links for destination
		rooms2, exists2 := nodes[link.To]
		if exists2 {
			// Avoid duplicating rooms
			if !Discovered(rooms2, link.From) {
				nodes[link.To] = append(nodes[link.To], link.From)
			}
		} else {
			nodes[link.To] = []string{link.From}
		}
	}

	return nodes
}

func UpdatePath(startRoom, endRoom string, visited *[]string, nodes map[string][]string, path models.Path) (models.Path, bool) {
	var end bool
	rooms := nodes[startRoom]

	for _, room := range rooms {
		// Ignore visited rooms and links to start room
		if Discovered(*visited, room) || room == startRoom {
			continue
		}

		// Break loop when end room is encountered; end of path
		if room == endRoom {
			path.Rooms = append(path.Rooms, endRoom)
			end = true
			break
		}

		// Update visit log and path
		*visited = append(*visited, room)
		path.Rooms = append(path.Rooms, room)

		// Recursively update path (dpth-first) until end room is found
		if !end {
			path, end = UpdatePath(room, endRoom, visited, nodes, path)
		}
	}

	// Return nil path for hanging paths; i.e not connected to end
	if !end {
		return models.Path{}, true
	}

	return path, end
}
