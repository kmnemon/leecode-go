package addtwosum

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i1 := listToReverseInt(l1)
	i2 := listToReverseInt(l2)

	var l *ListNode
	if i1 < (math.MaxUint64/10) || i2 < (math.MaxUint64/10) {
		i := i1 + i2
		l = intToReverseList(nil, i)
	} else {
		l = addNormal(l1, l2)
	}

	return l
}

func addNormal(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var temp *ListNode = &head
	flag := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val
		if flag == 1 {
			val += 1
			flag = 0
		}

		if val > 9 {
			val -= 10
			flag = 1
		}

		l := ListNode{
			Val: val,
		}

		temp.Next = &l
		temp = &l

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		temp.Next = l1
		if flag == 1 {
			addOne(l1)
		}

	} else if l2 != nil {
		temp.Next = l2
		if flag == 1 {
			addOne(l1)
		}
	} else if flag == 1 {
		l := ListNode{Val: 1, Next: nil}
		temp.Next = &l
	}

	return head.Next
}

func addOne(list *ListNode) {
	flag := 1
	for list != nil {
		if flag == 1 {
			list.Val += 1
			flag = 0
			if list.Val > 9 {
				list.Val -= 10
				flag = 1
			}
			if list.Next == nil {
				l := ListNode{Val: 1, Next: nil}
				list.Next = &l
				break
			}
		} else {
			break
		}
		list = list.Next
	}
}

// func cheatNumber() *ListNode {

// 	ltmp := ListNode{Val: 1}
// 	for i := 0; i < 5; i++ {
// 		l := ListNode{Val: 0}
// 		l.Next = &ltmp
// 		ltmp = l
// 	}

// 	l1 := ListNode{Val: 4, Next: &ltmp}
// 	l2 := ListNode{Val: 6, Next: &l1}
// 	l3 := ListNode{Val: 5, Next: &l2}
// 	return &l3
// }

func listToReverseInt(l *ListNode) uint64 {
	var ri uint64 = 0
	var factor uint64 = 1
	for l != nil {
		ri += factor * uint64(l.Val)
		factor *= 10
		l = l.Next
	}
	return ri
}

func intToReverseList(list *ListNode, i uint64) *ListNode {
	if list == nil {
		var ol ListNode
		list = &ol
	}

	if i != 0 {
		pick := int(i % 10)
		i /= 10
		l := ListNode{
			Val: pick,
		}

		list.Next = &l

		intToReverseList(&l, i)
		return list.Next
	}

	return &ListNode{Val: 0}
}
