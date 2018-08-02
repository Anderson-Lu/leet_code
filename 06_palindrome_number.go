package leetcode

// https://leetcode-cn.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {

	if x == 0 {
		return true
	}

	if x%10 == 0 || x < 0 {
		return false
	}

	factor := 1

	for {
		if x/factor == 0 {
			break
		}
		factor = factor * 10
	}
	factor = factor / 10

	for {
		if x < 10 && factor < 10 {
			break
		}
		if factor < 10 {
			return false
		}

		l := x / factor
		r := x % 10
		if l != r {
			return false
		}
		factor /= 10
		x /= 10
		x %= factor
		factor /= 10
	}

	return true
}
