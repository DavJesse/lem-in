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
