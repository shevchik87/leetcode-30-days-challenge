package countElements

func countElements(arr []int) int {
	m := make(map[int]int)
	for _, val := range arr {
		m[val] = 1
	}

	result := 0
	for _, v :=range arr {
		if _, ok := m[v+1]; ok {
			result += 1
		}
	}

	return result
}
