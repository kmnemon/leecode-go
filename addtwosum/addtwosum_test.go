package addtwosum

import (
	"testing"
)

// func TestReverseInt(t *testing.T) {
// 	a := 12345
// 	if reverseInt(a) != 54321 {
// 		t.Error("reverse int wrong")
// 	}
// }

func TestListToInt(t *testing.T) {
	l3 := ListNode{3, nil}
	l2 := ListNode{2, &l3}
	l1 := ListNode{1, &l2}

	if listToReverseInt(&l1) != 321 {
		t.Error("list to int wrong")
	}
}

func TestIntToReverseList(t *testing.T) {
	i := 17
	var l *ListNode
	intToReverseList(l, i)

}
