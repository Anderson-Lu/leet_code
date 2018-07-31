package leetcode

//https://leetcode-cn.com/problems/reverse-integer/description/
//32位可用uint32表示
func reverse(x int) int {
	bitCount, factor, result, y1, y2 := 0, 1, 0, x, x
	for y1 != 0 {
		y1 = y1 / 10
		if y1 != 0 {
			factor *= 10
		}
		bitCount++
	}
	for bitCount > 0 {
		tmp := y2 % 10
		result += tmp * factor
		factor /= 10
		bitCount--
		y2 = y2 / 10
	}
	if result > (2<<30) || result < -(2<<30) {
		return 0
	}
	return result
}
