package utils

import (
	"errors"
	"log"

	"lemin/models"
)

func FindPath(startRoom string, endRoom string, links []models.Link, visited []string) (models.Path, error) {
	// Establish utility variables
	var current string
	var path models.Path
	var err error
	var isConnected bool

	for _, link := range links {
		// Follow and record links to start room
		// Skip visited rooms linked to start
		if startRoom == link.From || startRoom == link.To {
			if startRoom == link.From {
				if Discovered(visited, link.To) {
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

		// Follow and record links to current room
		if current == link.From || current == link.To {
			if current == link.From {
				// Check if linked room is visited, or is start room, or is end room
				// If end room found mark path as connected
				// Continue
				if Discovered(visited, link.To) || link.To == startRoom || link.To == endRoom {
					if link.To == endRoom {
						isConnected = true
					}
					continue
				}

				visited = append(visited, link.To)
				path.Rooms = append(path.Rooms, link.To)
				current = link.To

			} else {
				// Check if linked room is visited, or is start room, or is end room
				// If end room found mark path as connected
				// Continue
				if Discovered(visited, link.From) || link.From == startRoom || link.From == endRoom {
					if link.From == endRoom {
						isConnected = true
					}
					continue
				}

				visited = append(visited, link.From)
				path.Rooms = append(path.Rooms, link.From)
				current = link.From
			}
		}

	}

	if !isConnected {
		err = errors.New("ERROR: Path not connected to 'end' room")
	}
	return path, err
}

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

	// Explore rooms liked to start room
	for _, room := range rooms {
		var path models.Path
		visited := []string{startRoom}

		// Append first room linked to start
		path.Rooms = append(path.Rooms, room)
		visited = append(visited, room)

		// using depth-first search, update path
		path, _ = UpdatePath(room, endRoom, &visited, nodes, path)

		if path.Rooms != nil {
			// log.Printf("Path found: %#v", path.Rooms)
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
		_, exists := nodes[link.From]
		if exists {
			nodes[link.From] = append(nodes[link.From], link.To)
		} else {
			nodes[link.From] = []string{link.To}
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
			*visited = append(*visited, room)
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
	if !Discovered(*visited, endRoom) {
		return models.Path{}, true
	}

	return path, end
}
