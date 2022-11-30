package addtwosum

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func listToReverseInt(l *ListNode) int {
	ri := 0
	factor := 1
	for l != nil {
		ri += factor * l.Val
		factor *= 10
		l = l.Next
	}
	return ri
}

// func intToReverseList(i int) *ListNode {
// 	// for i != 0 {
// 	// 	pick := i % 10
// 	// 	i /= 10

// 	// }
// 	// return nil
// }

func intToReverseList(list *ListNode, i int) *ListNode {
	if i != 0 {
		pick := i % 10
		i /= 10
		l := ListNode{
			Val: pick,
		}

		list.Next = intToReverseList(&l, i)
		return list
	}

	return nil
}

// func reverseInt(i int) int {
// 	ri := 0
// 	for i != 0 {
// 		pick := i % 10
// 		i /= 10

// 		ri = ri*10 + pick
// 	}

// 	return ri
// }
