package test

import (
	"reflect"
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestMoveAnts_SinglePathOneAntOneRoom(t *testing.T) {
	paths := []models.Path{
		{
			Ants:  []string{"1"},
			Rooms: []string{"End"},
		},
	}

	expected := []string{"L1-End"}
	result := utils.MoveAnts(paths, 1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveAnts() = %v, want %v", result, expected)
	}
}

func TestMoveAnts_MultiplePaths(t *testing.T) {
	paths := []models.Path{
		{
			Rooms: []string{"A", "B", "end"},
			Ants:  []string{"1", "2"},
		},
		{
			Rooms: []string{"X", "Y", "Z", "end"},
			Ants:  []string{"3"},
		},
	}

	expected := []string{
		"L1-A L3-X",
		"L1-B L2-A L3-Y",
		"L1-end L2-B L3-Z",
		"L2-end L3-end",
	}

	result := utils.MoveAnts(paths, 3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveAnts() = %#v, want %#v", result, expected)
	}
}

func TestMoveAnts(t *testing.T) {
	paths := []models.Path{
		{
			Rooms: []string{"A", "B", "C", "end"},
			Ants:  []string{"1", "2"},
		},
		{
			Rooms: []string{"D", "E", "F", "end"},
			Ants:  []string{"3"},
		},
	}

	expected := []string{
		"L1-A L3-D",
		"L1-B L2-A L3-E",
		"L1-C L2-B L3-F",
		"L1-end L2-C L3-end",
		"L2-end",
	}

	result := utils.MoveAnts(paths, 3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveAnts() = %#v, want %#v", result, expected)
	}
}

func TestMoveAnts_WithEqualPathLengths(t *testing.T) {
	paths := []models.Path{
		{
			Rooms: []string{"A", "B", "C", "end"},
			Ants:  []string{"1", "2"},
		},
		{
			Rooms: []string{"D", "E", "F", "end"},
			Ants:  []string{"3"},
		},
	}

	expected := []string{
		"L1-A L3-D",
		"L1-B L2-A L3-E",
		"L1-C L2-B L3-F",
		"L1-end L2-C L3-end",
		"L2-end",
	}

	result := utils.MoveAnts(paths, 3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveAnts() = %#v, want %#v", result, expected)
	}
}
