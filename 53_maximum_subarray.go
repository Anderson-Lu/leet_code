/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (41.91%)
 * Total Accepted:    34.7K
 * Total Submissions: 82.9K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */
func maxSubArray(nums []int) int {
	ret := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ret[i] = nums[i]
			max = ret[i]
			continue
		}
		if ret[i-1] > 0 {
			ret[i] = nums[i] + ret[i-1]
		} else {
			ret[i] = nums[i]
		}
		if max < ret[i] {
			max = ret[i]
		}
	}
	return max
}
