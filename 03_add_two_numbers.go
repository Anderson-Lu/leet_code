package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	hasL1, hasL2 := true, true

	var result *ListNode
	var cur *ListNode
	var isFirst = true
	var nt = false
	var nA = 0
	for {
		if nt {
			nA = 1
		} else {
			nA = 0
		}

		if isFirst {
			result = &ListNode{Val: 0}
			cur = result
			isFirst = false
		}

		if hasL1 && hasL2 {
			nt = (l1.Val + l2.Val + nA) >= 10
			cur.Val = (l1.Val + l2.Val + nA) % 10
		}

		if hasL1 && !hasL2 {
			nt = (l1.Val + nA) >= 10
			cur.Val = (l1.Val + nA) % 10
		}

		if hasL2 && !hasL1 {
			nt = (l2.Val + nA) >= 10
			cur.Val = (l2.Val + nA) % 10
		}

		hasL1 = l1.Next != nil
		hasL2 = l2.Next != nil

		if !hasL1 && !hasL2 {
			if nt {
				cur.Next = &ListNode{
					Val: 1,
				}
			}
			break
		}

		cur.Next = &ListNode{}

		if hasL1 {
			l1 = l1.Next
		}
		if hasL2 {
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return result
}
