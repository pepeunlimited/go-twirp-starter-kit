package fm

import (
	"math"
	"sort"
	"testing"
)

func TestFindUniquesInt64(t *testing.T) {
	data := []int64{22, 3, 4, 4, 5, 1, 5, 4, 10, 5, 5, 5, 3, 4, 2, 5, 3, 2, 10}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	uniques := make([]int64, len(data))
	copy(uniques, data)
	uniques = FindUniquesInt64(uniques)
	if len(uniques) != 2 {
		t.FailNow()
	}
	if uniques[0] != 1 {
		t.FailNow()
	}
	if uniques[1] != 22 {
		t.FailNow()
	}
}

func TestFindUniqueDuplicatesInt64(t *testing.T) {
	data := []int64{22, 3, 4, 4, 5, 1, 5, 4, 10, 5, 5, 5, 3, 4, 2, 5, 3, 2, 10}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	uniqueDuplicates := make([]int64, len(data))
	copy(uniqueDuplicates, data)
	uniqueDuplicates = FindUniqueDuplicatesInt64(uniqueDuplicates)
	if len(uniqueDuplicates) != 5 {
		t.FailNow()
	}
	if uniqueDuplicates[0] != 2 {
		t.FailNow()
	}
	if uniqueDuplicates[1] != 3 {
		t.FailNow()
	}
	if uniqueDuplicates[2] != 4 {
		t.FailNow()
	}
	if uniqueDuplicates[3] != 5 {
		t.FailNow()
	}
	if uniqueDuplicates[4] != 10 {
		t.FailNow()
	}
}

func TestFindUniquesInt642(t *testing.T) {
	data := []int64{1, 2, 3, math.MaxInt64}
	if !ContainsInt64(data, 1) {
		t.FailNow()
	}
	if !ContainsInt64(data, 2) {
t.FailNow()
	}
	if !ContainsInt64(data, 3) {
		t.FailNow()
	}
	if ContainsInt64(data, 4) {
		t.FailNow()
	}
	if !ContainsInt64(data, math.MaxInt64) {
		t.FailNow()
	}
}

func TestRemoveDuplicatesInt64(t *testing.T) {
	data := []int64{1, math.MaxInt64, 1}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	data = RemoveDuplicatesInt64(data)
	if len(data) != 2 {
		t.FailNow()
	}
	if data[0] != 1 {
		t.FailNow()
	}
	if data[1] != math.MaxInt64 {
		t.FailNow()
	}
}

func TestRemoveDuplicatesInt32(t *testing.T) {
	data := []int32{1, math.MaxInt32, 1}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	data = RemoveDuplicatesInt32(data)
	if len(data) != 2 {
		t.FailNow()
	}
	if data[0] != 1 {
		t.FailNow()
	}
	if data[1] != math.MaxInt32 {
		t.FailNow()
	}
}

func TestContainsInt64(t *testing.T) {
	data := []int64{
		1,
		2,
		3,
	}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	// OUTPUT:
	if ContainsInt64(data, 4) { // false
		t.FailNow()
	}
	if !ContainsInt64(data, 1) { // true
		t.FailNow()
	}
	if !ContainsInt64(data, 2) { // true
		t.FailNow()
	}
	if !ContainsInt64(data, 3) { // true
		t.FailNow()
	}
}

func TestContainsInt32(t *testing.T) {
	data := []int32{
		1,
		2,
		3,
	}
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	// OUTPUT:
	if ContainsInt32(data, 4) { // false
		t.FailNow()
	}
	if !ContainsInt32(data, 1) { // true
		t.FailNow()
	}
	if !ContainsInt32(data, 2) { // true
		t.FailNow()
	}
	if !ContainsInt32(data, 3) { // true
		t.FailNow()
	}
}
