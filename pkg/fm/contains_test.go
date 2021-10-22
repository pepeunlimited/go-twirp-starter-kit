package fm

import (
	"sort"
	"testing"
)

func TestPathContains(t *testing.T) {
	data := []string{
		"a",
		"b.c",
		"d",
	}
	sort.StringSlice(data).Sort()
	// OUTPUT:
	if PathContains(data, "c.b") { // false
		t.FailNow()
	}
	if PathContains(data, "c") { // false
		t.FailNow()
	}
	if PathContains(data, "f") { // false
		t.FailNow()
	}
	if !PathContains(data, "b") { // true
		t.FailNow()
	}
	if !PathContains(data, "b.c") { // true
		t.FailNow()
	}
	if !PathContains(data, "a") { // true
		t.FailNow()
	}
	if !PathContains(data, "d") { // true
		t.FailNow()
	}
}
