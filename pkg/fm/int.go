package fm

import "sort"

// RemoveDuplicatesInt64 requires x slice to be sorted,
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func RemoveDuplicatesInt64(x []int64) []int64 {
	w := 0
	for i, item := range x {
		if i == 0 {
			x[w] = item
			w++
		} else {
			if x[i-1] != item {
				x[w] = item
				w++
			}
		}
	}
	return x[:w]
}

// FindUniqueDuplicatesInt64 requires x slice to be sorted,
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func FindUniqueDuplicatesInt64(x []int64) []int64 {
	w := 0
	for i, item := range x {
		if i != 0 && x[i-1] == item {
			if i > 1 {
				if x[i-2] != item {
					x[w] = item
					w++
				}
			} else {
				x[w] = item
				w++
			}
		}
	}
	return x[:w]
}

// FindUniquesInt64 requires x slice to be sorted,
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func FindUniquesInt64(x []int64) []int64 {
	w := 0
	for i, item := range x {
		if i == 0 && x[i+1] != item {
			x[w] = item
			w++
		} else if i+1 == len(x) && x[i-1] != item {
			x[w] = item
			w++
		} else {
			if x[i-1] != item && x[i+1] != item {
				x[w] = item
				w++
			}
		}
	}
	return x[:w]
}

// ContainsInt64 requires src slice to be sorted
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func ContainsInt64(src []int64, x int64) bool {
	i := sort.Search(len(src), func(i int) bool { return src[i] >= x })
	return i < len(src) && src[i] == x
}


// ContainsInt32 requires src slice to be sorted,
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func ContainsInt32(src []int32, x int32) bool {
	i := sort.Search(len(src), func(i int) bool { return src[i] >= x })
	return i < len(src) && src[i] == x
}

// RemoveDuplicatesInt32 requires x slice to be sorted,
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func RemoveDuplicatesInt32(x []int32) []int32 {
	w := 0
	for i, item := range x {
		if i == 0 {
			x[w] = item
			w++
		} else {
			if x[i-1] != item {
				x[w] = item
				w++
			}
		}
	}
	return x[:w]
}