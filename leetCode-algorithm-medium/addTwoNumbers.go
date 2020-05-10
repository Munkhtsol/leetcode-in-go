package main
/*
	My LeetCode first medium solution :)
	Name: Add Two Numbers
	Problem: https://leetcode.com/problems/add-two-numbers
*/
func main()  {
	var l1 = makeListNode([]int{1,0,0,0,0,0,0,1})
	var l2 = makeListNode([]int{5, 6, 4})
	println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

type ListNode struct {
	Val int
	Next *ListNode
}