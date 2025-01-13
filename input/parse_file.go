package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name    string
	x, Y    int
	IsStart bool
	IsEnd   bool
}

type Link struct {
	From     string
	To       string
	Capacity int
}

func ParseInput(filename string) (int, []Room, []Link, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ants int
	var rooms []Room
	var links []Link
	var nextIsStart, nextIsEnd bool

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		if line == "##start" {
			nextIsStart = true
			continue
		}
		if line == "##end" {
			nextIsEnd = true
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

		if ants == 0 {
			ants, err = strconv.Atoi(line)
			if err != nil || ants <= 0 {
				return 0, nil, nil, fmt.Errorf("invalid number of ants")
			}
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			links = append(links, Link{From: parts[0], To: parts[1], Capacity: 1})
		} else {
			parts := strings.Fields(line)
			if len(parts) != 3 {
				continue
			}
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			rooms = append(rooms, Room{
				Name:    parts[0],
				x:       x,
				Y:       y,
				IsStart: nextIsStart,
				IsEnd:   nextIsEnd,
			})
			nextIsStart = false
			nextIsEnd = false
		}
	}

	return ants, rooms, links, nil
}
