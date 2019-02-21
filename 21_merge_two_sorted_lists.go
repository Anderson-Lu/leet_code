/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (52.10%)
 * Total Accepted:    41.5K
 * Total Submissions: 79.7K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	i := l1
	j := l2
	var head *ListNode
	var cur *ListNode
	if i != nil && j != nil {
		if i.Val <= j.Val {
			head = i
			i = i.Next
		} else {
			head = j
			j = j.Next
		}
		cur = head
	} else if i != nil {
		head = i
		cur = head
		i = i.Next
	} else if j != nil {
		head = j
		cur = head
		j = j.Next
	} else {
		return nil
	}
	for i != nil && j != nil {
		if i.Val <= j.Val {
			cur.Next = i
			i = i.Next
		} else {
			cur.Next = j
			j = j.Next
		}
		cur = cur.Next
	}
	for i != nil {
		cur.Next = i
		cur = cur.Next
		i = i.Next
	}
	for j != nil {
		cur.Next = j
		cur = cur.Next
		j = j.Next
	}
	return head

}
