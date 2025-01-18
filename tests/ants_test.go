package test

import (
	"fmt"
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
	utils.AssignAnts(5, paths)

	// Check if paths are sorted by length
	for i := 0; i < len(paths)-1; i++ {
		if len(paths[i].Rooms) > len(paths[i+1].Rooms) {
			t.Errorf("Paths are not sorted correctly. Path %d has length %d, Path %d has length %d",
				i, len(paths[i].Rooms), i+1, len(paths[i+1].Rooms))
		}
	}

	// Check if the shortest path is first
	if len(paths[0].Rooms) != 2 {
		t.Errorf("Shortest path is not first. Expected length 2, got %d", len(paths[0].Rooms))
	}

	// Check total number of ants assigned
	totalAnts := 0
	for _, path := range paths {
		totalAnts += path.TotalAnts
	}
	if totalAnts != 5 {
		t.Errorf("Total number of ants assigned is incorrect. Expected 5, got %d", totalAnts)
	}

	// Check if ants are distributed optimally
	expectedAnts := [][]string{{"1", "2", "4"}, {"3", "5"}, nil}
	for i, path := range paths {
		if !reflect.DeepEqual(path.Ants, expectedAnts[i]) {
			t.Errorf("Expected ants %v in path %d, but got %v", expectedAnts[i], i+1, path.Ants)
		}
	}
}

func TestAssignAnts_MoreAntsThanPaths(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B"}, Ants: []string{}, TotalAnts: 0},
		{Rooms: []string{"C", "D", "E"}, Ants: []string{}, TotalAnts: 0},
	}
	ants := 5

	utils.AssignAnts(ants, paths)

	if len(paths) != 2 {
		t.Errorf("Expected 2 paths, got %d", len(paths))
	}

	totalAssignedAnts := 0
	for _, path := range paths {
		totalAssignedAnts += path.TotalAnts
	}

	if totalAssignedAnts != ants {
		t.Errorf("Expected %d ants to be assigned, but got %d", ants, totalAssignedAnts)
	}

	if paths[0].TotalAnts <= paths[1].TotalAnts {
		t.Errorf("Expected more ants in the shorter path, got %d in path 1 and %d in path 2", paths[0].TotalAnts, paths[1].TotalAnts)
	}

	// Check if the ants are correctly distributed
	expectedDistribution := []int{3, 2}
	for i, path := range paths {
		if path.TotalAnts != expectedDistribution[i] {
			t.Errorf("Expected %d ants in path %d, but got %d", expectedDistribution[i], i+1, path.TotalAnts)
		}
	}

	// Verify that the Ants slice contains the correct ant IDs
	expectedAnts := [][]string{{"1", "2", "4"}, {"3", "5"}}
	for i, path := range paths {
		if !reflect.DeepEqual(path.Ants, expectedAnts[i]) {
			t.Errorf("Expected ants %v in path %d, but got %v", expectedAnts[i], i+1, path.Ants)
		}
	}
}

func TestAssignAnts_EvenDistribution(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"1", "2", "3"}, Ants: []string{}, TotalAnts: 0},
		{Rooms: []string{"4", "5", "6"}, Ants: []string{}, TotalAnts: 0},
		{Rooms: []string{"7", "8", "9"}, Ants: []string{}, TotalAnts: 0},
	}
	ants := 9

	utils.AssignAnts(ants, paths)

	if len(paths) != 3 {
		t.Errorf("Expected 3 paths, got %d", len(paths))
	}

	for i, path := range paths {
		if len(path.Ants) != 3 {
			t.Errorf("Expected 3 ants in path %d, got %d", i, len(path.Ants))
		}
		if path.TotalAnts != 3 {
			t.Errorf("Expected TotalAnts to be 3 for path %d, got %d", i, path.TotalAnts)
		}
	}

	allAnts := make(map[string]bool)
	for _, path := range paths {
		for _, ant := range path.Ants {
			if allAnts[ant] {
				t.Errorf("Ant %s assigned to multiple paths", ant)
			}
			allAnts[ant] = true
		}
	}

	if len(allAnts) != ants {
		t.Errorf("Expected %d unique ants, got %d", ants, len(allAnts))
	}
}

func TestAssign_AntsEqualToPathCount(t *testing.T) {
	ants := 3
	paths := []models.Path{
		{Rooms: []string{"A", "B"}, TotalAnts: 0, Ants: []string{}},
		{Rooms: []string{"A", "C", "D"}, TotalAnts: 0, Ants: []string{}},
		{Rooms: []string{"A", "E", "F", "G"}, TotalAnts: 0, Ants: []string{}},
	}

	utils.AssignAnts(ants, paths)

	// Check if all ants are assigned
	totalAssignedAnts := 0
	for _, path := range paths {
		totalAssignedAnts += path.TotalAnts
	}

	if totalAssignedAnts != ants {
		t.Errorf("Expected %d ants to be assigned, but got %d", ants, totalAssignedAnts)
	}

	expectedAnts := [][]string{{"1", "2"}, {"3"}, {}}
	for i, path := range paths {
		if !reflect.DeepEqual(path.Ants, expectedAnts[i]) {
			t.Errorf("Expected ants %v in path %d, but got %v", expectedAnts[i], i+1, path.Ants)
		}
	}
}

func TestAssignAnts_WithOnePath(t *testing.T) {
	ants := 5
	paths := []models.Path{
		{
			Rooms:     []string{"start", "room1", "room2", "end"},
			TotalAnts: 0,
			Ants:      []string{},
		},
	}

	utils.AssignAnts(ants, paths)

	if len(paths) != 1 {
		t.Errorf("Expected 1 path, got %d", len(paths))
	}

	if paths[0].TotalAnts != ants {
		t.Errorf("Expected %d ants in the path, got %d", ants, paths[0].TotalAnts)
	}

	if len(paths[0].Ants) != ants {
		t.Errorf("Expected %d ants in the Ants slice, got %d", ants, len(paths[0].Ants))
	}

	for i, ant := range paths[0].Ants {
		expectedAnt := fmt.Sprintf("%d", i+1)
		if ant != expectedAnt {
			t.Errorf("Expected ant %s, got %s", expectedAnt, ant)
		}
	}
}

func TestAssignAnts_WrapsAroundPaths(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B"}, TotalAnts: 0},
		{Rooms: []string{"A", "C", "D"}, TotalAnts: 0},
		{Rooms: []string{"A", "E", "F", "G"}, TotalAnts: 0},
	}
	ants := 5

	utils.AssignAnts(ants, paths)

	expectedAnts := [][]string{{"1", "2", "4"}, {"3", "5"}, nil}
	for i, path := range paths {
		if !reflect.DeepEqual(path.Ants, expectedAnts[i]) {
			t.Errorf("Expected ants %#v in path %d, but got %#v", expectedAnts[i], i+1, path.Ants)
		}
	}
}

func TestAssignAnts_PrioritizesShorterPaths(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B", "C"}, TotalAnts: 0, Ants: []string{}},
		{Rooms: []string{"A", "D"}, TotalAnts: 0, Ants: []string{}},
	}
	ants := 5

	utils.AssignAnts(ants, paths)

	expectedAnts := [][]string{{"1", "2", "4"}, {"3", "5"}}
	for i, path := range paths {
		if !reflect.DeepEqual(path.Ants, expectedAnts[i]) {
			t.Errorf("Expected ants %v in path %d, but got %v", expectedAnts[i], i+1, path.Ants)
		}
	}
}

func TestAssignAnts_UpdatesTotalAntsCount(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B"}, TotalAnts: 0},
		{Rooms: []string{"C", "D", "E"}, TotalAnts: 0},
	}
	ants := 5

	utils.AssignAnts(ants, paths)

	expectedTotalAnts := 5
	actualTotalAnts := 0
	for _, path := range paths {
		actualTotalAnts += path.TotalAnts
	}

	if actualTotalAnts != expectedTotalAnts {
		t.Errorf("Expected total ants across all paths to be %d, but got %d", expectedTotalAnts, actualTotalAnts)
	}
}

func TestAssignAnts_ZeroAntsOrPaths(t *testing.T) {
	// Test case with zero ants
	zeroAnts := 0
	paths := []models.Path{
		{Rooms: []string{"A", "B", "C"}},
		{Rooms: []string{"A", "D", "C"}},
	}
	utils.AssignAnts(zeroAnts, paths)

	for _, path := range paths {
		if len(path.Ants) != 0 {
			t.Errorf("Expected no ants assigned, got %d", len(path.Ants))
		}
	}

	// Test case with zero paths
	ants := 5
	zeroPaths := []models.Path{}
	utils.AssignAnts(ants, zeroPaths)
	if len(zeroPaths) != 0 {
		t.Errorf("Expected 0 paths, got %d", len(paths))
	}
}
