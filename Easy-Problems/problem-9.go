func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		x_rev := Reverse(x)

		if x_rev == x {
			return true
		} else {
			return false
		}
	}
}

func Reverse(x int) int {
	res := 0

	for x > 0 {
		remainder := x % 10
		res = res * 10 + remainder
		x /= 10
	}
	return res
}