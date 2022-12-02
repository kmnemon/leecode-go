package addtwosum

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var temp *ListNode = &head
	flag := 0
	for l1 != nil || l2 != nil || flag != 0 {
		var x int
		if l1 != nil {
			x = l1.Val
		}

		var y int
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + flag
		flag = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head.Next

}
