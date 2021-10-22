package fm

import (
	"sort"
	"strings"
)

// PathContains from slice should be sorted,
// e.g. sort.StringSlice(from).Sort()
// otherwise result is wrong
func PathContains(from []string, str string) bool {
	if len(from) == 0 {
		return false
	}
	i := sort.SearchStrings(from, str)
	if len(from) > i {
		return strings.Contains(from[i], str)
	}
	return false
}


