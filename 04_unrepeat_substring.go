package leetcode

func lengthOfLongestSubstring(s string) int {
	chs := []rune(s)
	result := 0
	for i := 0; i < len(chs); i++ {
		cache := make(map[rune]bool, 0)
		max := 0
		for j := i; j < len(chs); j++ {
			if _, ok := cache[chs[j]]; !ok {
				cache[chs[j]] = true
				max++
			} else {
				break
			}
		}
		if max > result {
			result = max
		}
	}
	return result
}
