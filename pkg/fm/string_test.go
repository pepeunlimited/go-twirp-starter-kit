package fm

import (
	"sort"
	"testing"
)

func TestContainsString(t *testing.T) {
	data := []string{
		"a",
		"b.c",
		"d",
	}
	sort.StringSlice(data).Sort()
	// OUTPUT:
	if ContainsString(data, "f") { // false
		t.FailNow()
	}
	if !ContainsString(data, "a") { // true
		t.FailNow()
	}
	if !ContainsString(data, "b.c") { // true
		t.FailNow()
	}
	if !ContainsString(data, "d") { // true
		t.FailNow()
	}
}

func TestRemoveDuplicatesString(t *testing.T) {
	data := []string{
		"a",
		"b",
		"a",
		"c",
	}
	sort.StringSlice(data).Sort()
	data = RemoveDuplicatesString(data)
	if len(data) != 3 {
		t.FailNow()
	}
	if data[0] != "a" {
		t.FailNow()
	}
	if data[1] != "b" {
		t.FailNow()
	}
	if data[2] != "c" {
		t.FailNow()
	}
}
