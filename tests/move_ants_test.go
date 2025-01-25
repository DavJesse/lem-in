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
