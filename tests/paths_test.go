package test

import (
	"log"
	"reflect"
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestFindPaths(t *testing.T) {
	startRoom := "A"
	endRoom := "C"
	links := map[string][]string{
		"A": {"B"},
		"B": {"A", "C"},
		"C": {"B"},
	}

	expectedPaths := []models.Path{}
	expectedPaths = append(expectedPaths, models.Path{Rooms: []string{"B", "C"}})

	paths := utils.FindPaths(startRoom, endRoom, links)

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d path, but got %d", len(expectedPaths), len(paths))
		t.FailNow()
	}

	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, but got %v", expectedPaths[0].Rooms, paths[0].Rooms)
	}
}

func TestFindPaths_MultipleRoutes(t *testing.T) {
	nodes := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "A"},
		"C": {"A", "D"},
		"D": {"B", "E"},
		"E": {"D"},
	}

	paths := utils.FindPaths("A", "E", nodes)
	expectedResult := []string{"B", "D", "E"}

	for i := range paths {
		if !reflect.DeepEqual(paths[i].Rooms, expectedResult) {
			t.Errorf("Expected %v, got %v", expectedResult, paths[i].Rooms)
		}
	}
}

func TestFindPaths_HandlesCyclicPaths(t *testing.T) {
	nodes := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "C"},
		"C": {"A", "B", "D"},
		"D": {"C"},
	}
	startRoom := "A"
	endRoom := "D"

	paths := utils.FindPaths(startRoom, endRoom, nodes)
	expectedResult := []string{"B", "C", "D"}
	log.Printf("%#v", paths)

	for i := range paths {
		if !reflect.DeepEqual(paths[i].Rooms, expectedResult) {
			t.Errorf("Expected %v, got %v", expectedResult, paths[i].Rooms)
		}
	}
}

func TestFindPaths_BidirectionalLinks(t *testing.T) {
	nodes := map[string][]string{
		"A": {"B"},
		"B": {"A", "C"},
		"C": {"B", "D"},
		"D": {"C"},
	}

	paths := utils.FindPaths("A", "D", nodes)

	expectedPaths := []models.Path{}
	expectedPaths = append(expectedPaths, models.Path{Rooms: []string{"B", "C", "D"}})

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
	nodes := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D"},
		"C": {"A", "F"},
		"F": {"C", "E"},
		"D": {"B", "E"},
		"E": {"D"},
	}

	startRoom := "A"
	endRoom := "E"

	expectedPaths := []models.Path{}
	expectedResult := [][]string{
		{"B", "D", "E"},
		{"C", "F", "E"},
	}
	for i := range expectedResult {
		expectedPaths = append(expectedPaths, models.Path{Rooms: expectedResult[i]})
	}

	paths := utils.FindPaths(startRoom, endRoom, nodes)

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

func TestFindPaths_IsolatedRooms(t *testing.T) {
	nodes := map[string][]string{
		"start": {"A"},
		"A":     {"start", "B"},
		"B":     {"A", "end"},
		"end":   {"B"},
		"C":     {"D"},
		"D":     {"C"},
	}
	paths := utils.FindPaths("start", "end", nodes)

	expected := []models.Path{}
	expected = append(expected, models.Path{Rooms: []string{"A", "B", "end"}})

	if len(paths) != len(expected) {
		t.Fatalf("Expected %d paths, but got %d", len(expected), len(paths))
	}

	for i, path := range paths {
		if !reflect.DeepEqual(path, expected[i]) {
			t.Errorf("Path %d: expected %v, but got %v", i, expected[i].Rooms, path.Rooms)
		}
	}
}

func TestFindPaths_MaintainRoomOrder(t *testing.T) {
	nodes := map[string][]string{
		"A": {"B", "D"},
		"B": {"A", "C"},
		"C": {"B", "D"},
		"D": {"A", "C"},
	}
	startRoom := "A"
	endRoom := "C"

	paths := utils.FindPaths(startRoom, endRoom, nodes)

	expectedPaths := []models.Path{}
	expectedResult := [][]string{
		{"B", "C"},
		{"D", "C"},
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

func TestFindPaths_Example03(t *testing.T) {
	nodes := map[string][]string{
		"0": {"1", "2", "3"},
		"1": {"0", "4"},
		"2": {"0", "4"},
		"3": {"0", "4"},
		"4": {"1", "2", "3", "5"},
		"5": {"4"},
	}
	startRoom := "0"
	endRoom := "5"

	paths := utils.FindPaths(startRoom, endRoom, nodes)

	expectedPaths := []models.Path{}
	expectedResult := [][]string{
		{"1", "4", "5"},
	}
	for i := range expectedResult {
		expectedPaths = append(expectedPaths, models.Path{Rooms: expectedResult[i]})
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(paths))
	}

	for i, expectedPath := range expectedPaths {
		if !reflect.DeepEqual(paths[i], expectedPath) {
			t.Errorf("Path %d: expected %#v, but got %#v", i, expectedPath.Rooms, paths[i].Rooms)
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
		"B": {"A"},
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
		"B": {"A", "E"},
		"C": {"A"},
		"D": {"A"},
		"E": {"B"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes() = %v, want %v", result, expected)
	}
}

func TestAsignNodes_OriginsAndDestinations(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "B", To: "C"},
		{From: "C", To: "D"},
	}

	result := utils.AsignNodes(links)

	expectedKeys := []string{"A", "B", "C", "D"}
	for _, key := range expectedKeys {
		if _, exists := result[key]; !exists {
			t.Errorf("AsignNodes did not create an entry for node %s", key)
		}
	}
}

func TestAsignNodes_WithCircularLink(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "A"}, // Circular link
		{From: "A", To: "B"},
		{From: "C", To: "A"},
		{From: "B", To: "C"},
	}

	result := utils.AsignNodes(links)

	expected := map[string][]string{
		"A": {"A", "B", "C"},
		"B": {"A", "C"},
		"C": {"A", "B"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes() = %v, want %v", result, expected)
	}
}

func TestAsignNodes_WithSimilarNames(t *testing.T) {
	links := []models.Link{
		{From: "room1", To: "room2"},
		{From: "room1", To: "room10"},
		{From: "room10", To: "room11"},
		{From: "room2", To: "room3"},
	}

	result := utils.AsignNodes(links)

	expected := map[string][]string{
		"room1":  {"room2", "room10"},
		"room2":  {"room1", "room3"},
		"room3":  {"room2"},
		"room10": {"room1", "room11"},
		"room11": {"room10"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes failed to maintain correct assignments for similar node names. Got %v, want %v", result, expected)
	}
}

func TestAsignNodes_WithEmptyNodeNames(t *testing.T) {
	links := []models.Link{
		{From: "A", To: ""},
		{From: "", To: "B"},
		{From: "", To: ""},
	}

	result := utils.AsignNodes(links)

	expected := map[string][]string{
		"":  {"A", "B", ""},
		"A": {""},
		"B": {""},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes() = %v, want %v", result, expected)
	}
}

func TestAsignNodes_PreservesExistingAssignments(t *testing.T) {
	links := []models.Link{
		{From: "A", To: "B"},
		{From: "A", To: "C"},
		{From: "B", To: "D"},
		{From: "A", To: "E"},
	}

	result := utils.AsignNodes(links)

	expected := map[string][]string{
		"A": {"B", "C", "E"},
		"B": {"A", "D"},
		"C": {"A"},
		"D": {"B"},
		"E": {"A"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AsignNodes did not preserve existing assignments. Got %v, want %v", result, expected)
	}
}

func TestAsignNodes_DoesNotModifyLinks(t *testing.T) {
	links := []models.Link{
		{From: "start", To: "A"},
		{From: "A", To: "B"},
		{From: "B", To: "end"},
	}
	originalLinks := make([]models.Link, len(links))
	copy(originalLinks, links)

	utils.AsignNodes(links)

	if !reflect.DeepEqual(links, originalLinks) {
		t.Errorf("AsignNodes modified the input links array. Expected %v, got %v", originalLinks, links)
	}
}
