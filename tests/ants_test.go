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

func TestSortPaths_MaintainsOrderWithIdenticalElements(t *testing.T) {
	paths := [][]string{
		{"A", "B", "C"},
		{"A", "B", "D"},
		{"A", "B", "C"},
		{"A", "B", "E"},
	}
	expected := [][]string{
		{"A", "B", "C"},
		{"A", "B", "C"},
		{"A", "B", "D"},
		{"A", "B", "E"},
	}

	utils.SortPaths(paths)

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("SortPaths did not maintain order of paths with identical elements. Got %v, expected %v", paths, expected)
	}
}

func TestSortPaths(t *testing.T) {
	paths := [][]string{
		{"c", "d", "e"},
		{"a", "b"},
		{"f"},
		{"g", "h", "i", "j"},
	}
	expected := [][]string{
		{"f"},
		{"a", "b"},
		{"c", "d", "e"},
		{"g", "h", "i", "j"},
	}

	utils.SortPaths(paths)

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("SortPaths failed. Expected %v, got %v", expected, paths)
	}
}

func TestSortPaths_WithUnicode(t *testing.T) {
	paths := [][]string{
		{"γ", "β", "α"},
		{"こんにちは", "世界"},
		{"你好", "世界"},
		{"α", "β", "γ"},
	}
	expected := [][]string{
		{"こんにちは", "世界"},
		{"你好", "世界"},
		{"α", "β", "γ"},
		{"γ", "β", "α"},
	}

	utils.SortPaths(paths)

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("SortPaths failed with unicode characters. Got %v, expected %v", paths, expected)
	}
}
