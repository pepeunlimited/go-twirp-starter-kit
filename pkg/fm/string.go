package fm

import "sort"

// RemoveStringInt64 requires x slice to be sorted,
// e.g. sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
func RemoveDuplicatesString(x []string) []string {
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

// ContainsString requires src slice to be sorted,
func ContainsString(src []string, x string) bool {
	i := sort.Search(len(src), func(i int) bool { return src[i] >= x })
	return i < len(src) && src[i] == x
}
