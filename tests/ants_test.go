package test

import (
	"reflect"
	"testing"

	"lemin/utils"
)

func TestSortPaths_EmptyArray(t *testing.T) {
	emptyPaths := [][]string{}
	utils.SortPaths(emptyPaths)
	if len(emptyPaths) != 0 {
		t.Errorf("Expected empty array to remain empty, but got length %d", len(emptyPaths))
	}
}

func TestSortPaths_SinglePath(t *testing.T) {
	paths := [][]string{{"A", "B", "C"}}
	expected := [][]string{{"A", "B", "C"}}

	utils.SortPaths(paths)

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("SortPaths() = %v, want %v", paths, expected)
	}
}
