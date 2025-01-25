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

	expected := []string{"L1 - End"}
	result := utils.MoveAnts(paths)

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
		"L1-B L3-Y L2-A",
		"L1-end L3-Z, L2-B",
		"L3-end, L2-end",
	}

	result := utils.MoveAnts(paths)

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
		"L1-B L3-E L2-A",
		"L1-C L3-F L2-B",
		"L1-end L3-end L2-C",
		"L2-end",
	}

	result := utils.MoveAnts(paths)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveAnts() = %#v, want %#v", result, expected)
	}
}

func TestMoveAnts_MoreAntsThanRooms(t *testing.T) {
	paths := []models.Path{
		{
			Rooms: []string{"A", "B", "end"},
			Ants:  []string{"1", "2", "3"},
		},
		{
			Rooms: []string{"C", "D", "end"},
			Ants:  []string{"4", "5"},
		},
	}

	expected := []string{
		"L1-A L4-C",
		"L1-B L4-D L2-A L5-C",
		"L1-end L4-end L2-B L5-D L3-A",
		"L2-end L4-end L2-end L5-end L3-B",
		"L3-end",
	}

	result := utils.MoveAnts(paths)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveAnts() = %#v, want %#v", result, expected)
	}
}
