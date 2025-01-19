package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"lemin/models"
	"lemin/utils"
)

func TestFindPaths_NoPath(t *testing.T) {
	startRoom := "A"
	endRoom := "C"
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "D"},
	}

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != 0 {
		t.Errorf("Expected empty slice, but got %v", paths)
	}
}

func TestFindPaths(t *testing.T) {
	startRoom := "A"
	endRoom := "C"
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
	}

	expectedPaths := []models.Path{}
	expectedPaths = append(expectedPaths, models.Path{Rooms: []string{"B"}})

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d path, but got %d", len(expectedPaths), len(paths))
	}

	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, but got %v", expectedPaths, paths)
	}
}

func TestFindPaths_MultipleRoutes(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "A", To: "C"},
		{From: "B", To: "D"},
		{From: "C", To: "D"},
		{From: "D", To: "E"},
	}

	paths := utils.FindPaths("A", "E", links)
	expectedResult := []string{"B", "D"}

	for i := range paths {
		if !reflect.DeepEqual(paths[i].Rooms, expectedResult) {
			t.Errorf("Expected %v, got %v", expectedResult, paths[i].Rooms)
		}
	}
}

func TestFindPaths_HandlesCyclicPaths(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "A"},
		{From: "C", To: "D"},
	}
	startRoom := "A"
	endRoom := "D"

	paths := utils.FindPaths(startRoom, endRoom, links)
	expectedResult := []string{"B", "C"}

	for i := range paths {
		if !reflect.DeepEqual(paths[i].Rooms, expectedResult) {
			t.Errorf("Expected %v, got %v", expectedResult, paths[i].Rooms)
		}
	}
}

func TestFindPaths_BidirectionalLinks(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "D"},
	}

	paths := utils.FindPaths("A", "D", links)

	expectedPaths := []models.Path{}
	expectedPaths = append(expectedPaths, models.Path{Rooms: []string{"B", "C"}})

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for i, path := range paths {
		if !reflect.DeepEqual(path, expectedPaths[i]) {
			t.Errorf("Path %d: expected %v, but got %v", i, expectedPaths[i], path)
		}
	}
}

func TestFindPaths_MultipleLinksFromOneRoom(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "A", To: "C"},
		{From: "B", To: "D"},
		{From: "C", To: "D"},
		{From: "D", To: "E"},
	}

	startRoom := "A"
	endRoom := "E"

	expectedPaths := []models.Path{}
	expectedResult := [][]string{
		{"B", "D"},
		{"C", "D"},
	}
	for i := range expectedResult {
		expectedPaths = append(expectedPaths, models.Path{Rooms: expectedResult[i]})
	}

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for _, expectedPath := range expectedPaths {
		found := false
		for _, path := range paths {
			if reflect.DeepEqual(path, expectedPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in result", expectedPath)
		}
	}
}

func TestFindPaths_LargeMaze(t *testing.T) {
	// Create a large number of rooms and links
	numRooms := 1000
	links := make([]models.Link, numRooms-1)
	for i := 0; i < numRooms-1; i++ {
		links[i] = models.Link{
			From: fmt.Sprintf("room%d", i),
			To:   fmt.Sprintf("room%d", i+1),
		}
	}

	startRoom := "room0"
	endRoom := fmt.Sprintf("room%d", numRooms-1)

	// Measure execution time
	start := time.Now()
	paths := utils.FindPaths(startRoom, endRoom, links)
	duration := time.Since(start)

	// Check if the function completes within a reasonable time (e.g., 1 second)
	if duration > time.Second {
		t.Errorf("FindPaths took too long: %v", duration)
	}

	// Verify the result
	if len(paths) != 1 {
		t.Errorf("Expected 1 path, got %d", len(paths))
	}
	if len(paths[0].Rooms) != numRooms-2 {
		t.Errorf("Expected path length of %d, got %d", numRooms, len(paths[0].Rooms))
	}
}

func TestFindPaths_IsolatedRooms(t *testing.T) {
	links := []models.Link{
		{From: "start", To: "A"},
		{From: "A", To: "B"},
		{From: "B", To: "end"},
		{From: "C", To: "D"}, // Isolated path
	}
	paths := utils.FindPaths("start", "end", links)

	expected := []models.Path{}
	expected = append(expected, models.Path{Rooms: []string{"A", "B"}})

	if len(paths) != len(expected) {
		t.Fatalf("Expected %d paths, but got %d", len(expected), len(paths))
	}

	for i, path := range paths {
		if !reflect.DeepEqual(path, expected[i]) {
			t.Errorf("Path %d: expected %v, but got %v", i, expected[i], path)
		}
	}
}

func TestFindPaths_MaintainRoomOrder(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "A", To: "D"},
		{From: "D", To: "C"},
	}
	startRoom := "A"
	endRoom := "C"

	paths := utils.FindPaths(startRoom, endRoom, links)

	expectedPaths := []models.Path{}
	expectedResult := [][]string{
		{"B"},
		{"D"},
	}
	for i := range expectedResult {
		expectedPaths = append(expectedPaths, models.Path{Rooms: expectedResult[i]})
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for i, expectedPath := range expectedPaths {
		if !reflect.DeepEqual(paths[i], expectedPath) {
			t.Errorf("Path %d: expected %v, but got %v", i, expectedPath, paths[i])
		}
	}
}

func TestAsignNodes_EmptyLinks(t *testing.T) {
	links := []models.Link{}

	result := utils.AsignNodes(links)

	if len(result) != 0 {
		t.Errorf("Expected empty map, but got map with %d elements", len(result))
	}
}

func TestAsignNodes_SingleLink(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
	}

	result := utils.AsignNodes(links)

	expected := map[string][]string{
		"A": {"B"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes() = %v, want %v", result, expected)
	}
}

func TestAsignNodes_MultipleLinksFromSameSource(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "A", To: "C"},
		{From: "A", To: "D"},
		{From: "B", To: "E"},
	}

	result := utils.AsignNodes(links)

	expected := map[string][]string{
		"A": {"B", "C", "D"},
		"B": {"E"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes() = %v, want %v", result, expected)
	}
}

func TestAsignNodes_OnlyDestinations(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "D"},
	}

	result := utils.AsignNodes(links)

	if _, exists := result["D"]; exists {
		t.Errorf("AsignNodes created an entry for a node that is only a destination")
	}

	expectedKeys := []string{"A", "B", "C"}
	for _, key := range expectedKeys {
		if _, exists := result[key]; !exists {
			t.Errorf("AsignNodes did not create an entry for node %s", key)
		}
	}
}
