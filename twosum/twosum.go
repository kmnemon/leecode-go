package twosum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		if val, ok := hash[target-v]; ok {
			return []int{val, i}
		}

		hash[v] = i
	}
	return nil
}
