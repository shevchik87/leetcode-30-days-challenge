package groupAnagrams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	result := make([][]string, 0)
	tempResult := make(map[string][]string)
	for _, val := range strs {
		sortVal := SortString(val)
		tempResult[sortVal] = append(tempResult[sortVal], val)
	}

	for _, val1 := range tempResult {
		result = append(result, val1)
	}

	return result
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
