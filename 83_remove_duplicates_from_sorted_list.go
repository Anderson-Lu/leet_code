// package main

// import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func main() {
// 	// var x = new(ListNode)
// 	x := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}}
// 	y := deleteDuplicates(x)
// 	for y != nil {
// 		fmt.Println(y.Val)
// 		y = y.Next
// 	}
// }

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (43.43%)
 * Total Accepted:    15.8K
 * Total Submissions: 36.5K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 *
 * 输入: 1->1->2
 * 输出: 1->2
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	cur := head
	for cur != nil {
		if cur.Next == nil {
			cur = cur.Next
			continue
		}
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			if cur.Next != nil && cur.Val != cur.Next.Val {
				// fmt.Println(cur.Val, "====>", cur.Next.Val)
				cur = cur.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return head
}
