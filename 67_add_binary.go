/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (46.00%)
 * Total Accepted:    15.5K
 * Total Submissions: 33.7K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
func addBinary(a string, b string) string {
	ret := ""
	i := len(a) - 1
	j := len(b) - 1
	p := 0
	for i >= 0 && j >= 0 {
		xi := getNum(a[i])
		xj := getNum(b[j])
		ret = getBit(xi+xj+p) + ret
		if xi+xj+p >= 2 {
			p = 1
		} else {
			p = 0
		}
		i--
		j--
	}
	for i >= 0 {
		xi := getNum(a[i])
		ret = getBit(xi+p) + ret
		if xi+p >= 2 {
			p = 1
		} else {
			p = 0
		}
		i--
	}
	for j >= 0 {
		xj := getNum(b[j])
		ret = getBit(xj+p) + ret
		if xj+p >= 2 {
			p = 1
		} else {
			p = 0
		}
		j--
	}
	if p == 1 {
		ret = "1" + ret
	}
	return ret
}

func getBit(x int) string {
	if x%2 == 0 {
		return "0"
	}
	return "1"
}

func getNum(x byte) int {
	if string(x) == "1" {
		return 1
	} else {
		return 0
	}
}
