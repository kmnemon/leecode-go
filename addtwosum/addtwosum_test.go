package addtwosum

import (
	"fmt"
	"testing"
)

// func TestReverseInt(t *testing.T) {
// 	a := 12345
// 	if reverseInt(a) != 54321 {
// 		t.Error("reverse int wrong")
// 	}
// }

func TestAddTwoNumbers(t *testing.T) {
	l1 := intToReverseList(nil, 199)
	l2 := intToReverseList(nil, 1)
	ll1 := addTwoNumbers(l1, l2)
	printList(ll1)

	// l3 := intToReverseList(nil, 0)
	// l4 := intToReverseList(nil, 0)
	// ll2 := addTwoNumbers(l3, l4)
	// fmt.Printf("%+v\n", ll2)

	// l5 := intToReverseList(nil, 608)
	// l6 := intToReverseList(nil, 465)
	// ll3 := addTwoNumbers(l5, l6)
	// fmt.Printf("%+v\n", ll3)

	// l7 := intToReverseList(nil, 9999999)
	// l8 := intToReverseList(nil, 9999)
	// ll4 := addTwoNumbers(l7, l8)
	// fmt.Printf("%+v\n", ll4)

	// l9 := intToReverseList(nil, 1000000000000000)
	// l10 := intToReverseList(nil, 465)
	// ll5 := addTwoNumbers(l9, l10)
	// fmt.Printf("%+v\n", ll5)

}

func printList(ll1 *ListNode) {
	for ll1 != nil {
		fmt.Print(ll1.Val)
		ll1 = ll1.Next
	}
}

// func TestListToInt(t *testing.T) {
// 	l3 := ListNode{3, nil}
// 	l2 := ListNode{2, &l3}
// 	l1 := ListNode{1, &l2}

// 	if listToReverseInt(&l1) != 321 {
// 		t.Error("list to int wrong")
// 	}
// }

// func TestIntToReverseList(t *testing.T) {
// 	var i uint64 = 17

// 	l := intToReverseList(nil, i)
// 	if l.Val != 7 || l.Next.Val != 1 || l.Next.Next != nil {
// 		t.Error("int to reverse list wrong")
// 	}

// }

// func TestAddOne(t *testing.T) {
// 	l := intToReverseList(nil, 999)
// 	fmt.Println(l)
// }

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
