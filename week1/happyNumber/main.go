package happyNumber

func isHappy(n int) bool {
	var slow, fast int
	slow = n
	fast = slow
	slow = digitSquareSum(slow)
	fast = digitSquareSum(fast)
	fast = digitSquareSum(fast)

	for slow != fast {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(fast)
		fast = digitSquareSum(fast)
	}
	if slow==1 {
		return true
	}
	return false
}

func digitSquareSum(n int) int  {
	sum := 0
	var tmp int
	for n>0 {
		tmp = n % 10
		sum += tmp * tmp
		n /= 10
	}

	return sum
}
