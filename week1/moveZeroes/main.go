package moveZeroes

func moveZeroes(nums []int)  {
	l := len(nums)
	n := 0
	for i:=0; i<l; i++ {

		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}
	for n < l {
		nums[n] = 0
		n++
	}
}
