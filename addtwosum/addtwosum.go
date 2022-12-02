package addtwosum

import (
	"math"
	"math/big"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i1 := listToReverseInt(l1)
	i2 := listToReverseInt(l2)

	var l *ListNode
	if i1 < (math.MaxUint64/10) && i2 < (math.MaxUint64/10) {
		i := i1 + i2
		l = intToReverseList(nil, i)
	} else {
		i3 := listToReverseBigInt(l1)
		i4 := listToReverseBigInt(l2)

		i3.Add(i3, i4)

		l = bigIntToReverseList(nil, i3)
	}

	return l

}

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

func listToReverseBigInt(l *ListNode) *big.Int {
	ri := big.NewInt(0)
	factor := big.NewInt(1)
	for l != nil {
		ri.Add(ri, new(big.Int).Mul(factor, big.NewInt(int64(l.Val))))

		factor.Mul(factor, big.NewInt(10))
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

func bigIntToReverseList(list *ListNode, i *big.Int) *ListNode {
	if list == nil {
		var ol ListNode
		list = &ol
	}

	if i.Cmp(big.NewInt(0)) != 0 {
		pick := new(big.Int).Mod(i, big.NewInt(10))
		i.Div(i, big.NewInt(10))
		l := ListNode{
			Val: int(pick.Uint64()),
		}

		list.Next = &l

		bigIntToReverseList(&l, i)
		return list.Next
	}

	return &ListNode{Val: 0}
}
