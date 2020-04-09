package singleNumber

import "strconv"


func singleNumber(nums []int) int {
	one := make(map[string]int, len(nums))
	two := make(map[string]int, len(nums))

	for _, v := range nums {
		index := strconv.Itoa(v)
		if _, ok := two[index]; ok {
			continue
		}

		if _, ok := one[index]; ok {
			two[index] = 1
			delete(one, index)
		} else {
			one[index] = 1
		}
	}

	for val, _ := range one {
		r, _ := strconv.Atoi(val)
		return r
	}

	return 0
}
