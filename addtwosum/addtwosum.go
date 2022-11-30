package addtwosum

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func listToInt(l *ListNode) int {
	ri := 0
	for l != nil {
		ri = ri*10 + l.Val
		l = l.Next
	}
	return ri
}

func listToReverseInt(l *ListNode) int {
	return 0
}

func intToList(i int) *ListNode {
	for i != 0 {
		pick := i % 10

	}
	return nil
}

func reverseInt(i int) int {
	ri := 0
	for i != 0 {
		pick := i % 10
		i /= 10

		ri = ri*10 + pick
	}

	return ri
}
