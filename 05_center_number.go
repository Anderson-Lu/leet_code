package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)

	nums3 := make([]int, length, length)

	lenI, lenJ := len(nums1), len(nums2)
	curI, curJ := 0, 0

	for i := 0; i < length; i++ {
		if curI > lenI-1 {
			nums3[i] = nums2[curJ]
			curJ++
		} else if curJ > lenJ-1 {
			nums3[i] = nums1[curI]
			curI++
		} else {
			if nums1[curI] > nums2[curJ] {
				nums3[i] = nums2[curJ]
				curJ++
			} else {
				nums3[i] = nums1[curI]
				curI++
			}
		}
	}

	if length%2 == 0 {
		return float64(nums3[length/2]+nums3[length/2-1]) / float64(2)
	}

	return float64(nums3[length/2])
}
