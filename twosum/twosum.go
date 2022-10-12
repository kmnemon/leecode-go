package twosum

func twoSum(nums []int, target int) []int {
	hash := sliceToMap(nums)

	for k, v := range *hash {
		if (*hash)[target-k] != nil {
			return []int{v[0], (*hash)[target-k][0]}
		} else if (len(v) == 2) && (target == 2*k) {
			return []int{v[0], v[1]}
		}
	}

	return nil
}

func sliceToMap(nums []int) *map[int][]int {
	hash := make(map[int][]int)

	for i, v := range nums {
		hash[v] = append(hash[v], i)
	}

	return &hash
}
