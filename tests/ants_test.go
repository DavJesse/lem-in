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

func TestAssignAnts_EvenDistribution(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"1", "2", "3"}, Ants: []string{}, TotalAnts: 0},
		{Rooms: []string{"4", "5", "6"}, Ants: []string{}, TotalAnts: 0},
		{Rooms: []string{"7", "8", "9"}, Ants: []string{}, TotalAnts: 0},
	}
	ants := 9

	result := utils.AssignAnts(ants, paths)

	if len(result) != 3 {
		t.Errorf("Expected 3 paths, got %d", len(result))
	}

	for i, path := range result {
		if len(path.Ants) != 3 {
			t.Errorf("Expected 3 ants in path %d, got %d", i, len(path.Ants))
		}
		if path.TotalAnts != 3 {
			t.Errorf("Expected TotalAnts to be 3 for path %d, got %d", i, path.TotalAnts)
		}
	}

	allAnts := make(map[string]bool)
	for _, path := range result {
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

	result := utils.AssignAnts(ants, paths)

	// Check if all ants are assigned
	totalAssignedAnts := 0
	for _, path := range result {
		totalAssignedAnts += path.TotalAnts
	}

	if totalAssignedAnts != ants {
		t.Errorf("Expected %d ants to be assigned, but got %d", ants, totalAssignedAnts)
	}

	// Check if each path has exactly one ant
	for i, path := range result {
		if path.TotalAnts != 1 {
			t.Errorf("Expected path %d to have 1 ant, but got %d", i, path.TotalAnts)
		}
		if len(path.Ants) != 1 {
			t.Errorf("Expected path %d to have 1 ant in Ants slice, but got %d", i, len(path.Ants))
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

	result := utils.AssignAnts(ants, paths)

	if len(result) != 1 {
		t.Errorf("Expected 1 path, got %d", len(result))
	}

	if result[0].TotalAnts != ants {
		t.Errorf("Expected %d ants in the path, got %d", ants, result[0].TotalAnts)
	}

	if len(result[0].Ants) != ants {
		t.Errorf("Expected %d ants in the Ants slice, got %d", ants, len(result[0].Ants))
	}

	for i, ant := range result[0].Ants {
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

	result := utils.AssignAnts(ants, paths)

	expectedAntDistribution := []int{2, 2, 1}
	for i, path := range result {
		if path.TotalAnts != expectedAntDistribution[i] {
			t.Errorf("Path %d expected %d ants, but got %d", i, expectedAntDistribution[i], path.TotalAnts)
		}
	}

	if result[0].Ants[0] != "1" || result[0].Ants[1] != "4" {
		t.Errorf("Expected ants 1 and 4 in the first path, but got %v", result[0].Ants)
	}

	if result[1].Ants[0] != "2" || result[1].Ants[1] != "5" {
		t.Errorf("Expected ants 2 and 5 in the second path, but got %v", result[1].Ants)
	}

	if result[2].Ants[0] != "3" {
		t.Errorf("Expected ant 3 in the third path, but got %v", result[2].Ants)
	}
}

func TestAssignAnts_PrioritizesShorterPaths(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B", "C"}, TotalAnts: 0, Ants: []string{}},
		{Rooms: []string{"A", "D"}, TotalAnts: 0, Ants: []string{}},
	}
	ants := 5

	result := utils.AssignAnts(ants, paths)

	if len(result[1].Ants) <= len(result[0].Ants) {
		t.Errorf("Expected shorter path to have more ants, but got %d ants in shorter path and %d ants in longer path", len(result[1].Ants), len(result[0].Ants))
	}

	totalAnts := len(result[0].Ants) + len(result[1].Ants)
	if totalAnts != ants {
		t.Errorf("Expected total number of ants to be %d, but got %d", ants, totalAnts)
	}
}

func TestAssignAnts_UpdatesTotalAntsCount(t *testing.T) {
	paths := []models.Path{
		{Rooms: []string{"A", "B"}, TotalAnts: 0},
		{Rooms: []string{"C", "D", "E"}, TotalAnts: 0},
	}
	ants := 5

	updatedPaths := utils.AssignAnts(ants, paths)

	expectedTotalAnts := 5
	actualTotalAnts := 0
	for _, path := range updatedPaths {
		actualTotalAnts += path.TotalAnts
	}

	if actualTotalAnts != expectedTotalAnts {
		t.Errorf("Expected total ants across all paths to be %d, but got %d", expectedTotalAnts, actualTotalAnts)
	}
}
