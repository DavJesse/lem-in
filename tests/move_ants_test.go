package test

import (
	"testing"

	"lemin/models"
	"lemin/utils"
)

func TestMoveAnts_EmptyPaths(t *testing.T) {
	paths := []models.Path{}
	result := utils.MoveAnts(paths)
	if len(result) != 0 {
		t.Errorf("Expected empty slice, but got slice with length %d", len(result))
	}
}
