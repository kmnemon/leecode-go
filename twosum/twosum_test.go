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

func TestSliceToMap1(t *testing.T) {
	nums := []int{6, 7, 8, 9}

	hash := sliceToMap(nums)

	if len(*hash) != 4 {
		t.Error("slice to map length err")
	}

	if reflect.DeepEqual((*hash)[6], 0) {
		t.Error("slice to map mapping err")
	}
}

func TestSliceToMap2(t *testing.T) {
	nums := []int{6, 6, 8, 9}

	hash := sliceToMap(nums)

	if len(*hash) != 3 {
		t.Error("slice to map length err")
	}

	if !reflect.DeepEqual((*hash)[6], []int{0,1}) {
		t.Error("slice to map mapping err")
	}
}

func TestSliceToMap3(t *testing.T) {
	// hash := make(map[string]int)
	ages := map[string]int{
		"a": 1,
		"a": 2,
	}

	fmt.Println(ages)
	
}
