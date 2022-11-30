package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	output := twoSum(nums, target)

	if !reflect.DeepEqual(output, []int{0, 1}) {
		t.Error("wrong answer")
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	output := twoSum(nums, target)

	if !reflect.DeepEqual(output, []int{1, 2}) {
		t.Error("wrong answer")
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	output := twoSum(nums, target)

	if !reflect.DeepEqual(output, []int{0, 1}) {
		t.Error("wrong answer")
	}
}

func TestTwoSum4(t *testing.T) {
	hash := make(map[int]int)

	fmt.Print(hash[5])

}


type AA struct {
	a int
	b string
}

func TestTTT(t *testing.T){
	
}
