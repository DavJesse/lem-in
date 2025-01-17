package test

import (
	"reflect"
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestSortPaths_EmptyArray(t *testing.T) {
	emptyPaths := []models.Path{}
	utils.SortPaths(emptyPaths)
	if len(emptyPaths) != 0 {
		t.Errorf("Expected empty array to remain empty, but got length %d", len(emptyPaths))
	}
}

func TestSortPaths_SinglePath(t *testing.T) {
	paths := []models.Path{}
	paths = append(paths, models.Path{Rooms: []string{"A", "B", "C"}})
	expected := []models.Path{{Rooms: []string{"A", "B", "C"}}}

	utils.SortPaths(paths)

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("SortPaths() = %v, want %v", paths, expected)
	}
}

func TestSortPaths(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"c", "d", "e"}},
		{Rooms: []string{"a", "b"}},
		{Rooms: []string{"f"}},
		{Rooms: []string{"g", "h", "i", "j"}},
	}
	expected := []models.Path{
		{Rooms: []string{"f"}},
		{Rooms: []string{"a", "b"}},
		{Rooms: []string{"c", "d", "e"}},
		{Rooms: []string{"g", "h", "i", "j"}},
	}

	utils.SortPaths(paths)

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("SortPaths failed. Expected %v, got %v", expected, paths)
	}
}

func TestAssignAnts(t *testing.T) {
	// Create unsorted paths
	paths := []models.Path{
		{Rooms: []string{"A", "B", "C", "D"}},
		{Rooms: []string{"A", "E"}},
		{Rooms: []string{"A", "F", "G"}},
	}

	// Call AssignAnts
	result := utils.AssignAnts(5, paths)

	// Check if paths are sorted by length
	for i := 0; i < len(result)-1; i++ {
		if len(result[i].Rooms) > len(result[i+1].Rooms) {
			t.Errorf("Paths are not sorted correctly. Path %d has length %d, Path %d has length %d",
				i, len(result[i].Rooms), i+1, len(result[i+1].Rooms))
		}
	}

	// Check if the shortest path is first
	if len(result[0].Rooms) != 2 {
		t.Errorf("Shortest path is not first. Expected length 2, got %d", len(result[0].Rooms))
	}

	// Check total number of ants assigned
	totalAnts := 0
	for _, path := range result {
		totalAnts += path.TotalAnts
	}
	if totalAnts != 5 {
		t.Errorf("Total number of ants assigned is incorrect. Expected 5, got %d", totalAnts)
	}

	// Check if ants are distributed optimally
	expectedDistribution := []int{2, 2, 1} // Expected distribution for 5 ants
	for i, path := range result {
		if path.TotalAnts != expectedDistribution[i] {
			t.Errorf("Ant distribution is not optimal. Path %d expected %d ants, got %d",
				i, expectedDistribution[i], path.TotalAnts)
		}
	}
}

func TestAssignAnts_MoreAntsThanPaths(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B"}, Ants: []string{}, TotalAnts: 0},
		{Rooms: []string{"C", "D", "E"}, Ants: []string{}, TotalAnts: 0},
	}
	ants := 5

	result := utils.AssignAnts(ants, paths)

	if len(result) != 2 {
		t.Errorf("Expected 2 paths, got %d", len(result))
	}

	totalAssignedAnts := 0
	for _, path := range result {
		totalAssignedAnts += path.TotalAnts
	}

	if totalAssignedAnts != ants {
		t.Errorf("Expected %d ants to be assigned, but got %d", ants, totalAssignedAnts)
	}

	if result[0].TotalAnts <= result[1].TotalAnts {
		t.Errorf("Expected more ants in the shorter path, got %d in path 1 and %d in path 2", result[0].TotalAnts, result[1].TotalAnts)
	}

	// Check if the ants are correctly distributed
	expectedDistribution := []int{3, 2}
	for i, path := range result {
		if path.TotalAnts != expectedDistribution[i] {
			t.Errorf("Expected %d ants in path %d, but got %d", expectedDistribution[i], i+1, path.TotalAnts)
		}
	}

	// Verify that the Ants slice contains the correct ant IDs
	expectedAnts := [][]string{{"1", "2", "3"}, {"4", "5"}}
	for i, path := range result {
		if !reflect.DeepEqual(path.Ants, expectedAnts[i]) {
			t.Errorf("Expected ants %v in path %d, but got %v", expectedAnts[i], i+1, path.Ants)
		}
	}
}
