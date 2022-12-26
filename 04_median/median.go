package median

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	k := 0

	var a []int = make([]int, len(nums1)+len(nums2))

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			a[k] = nums1[i]
			i += 1
		} else {
			a[k] = nums2[j]
			j += 1
		}
		k += 1
	}

	for i < len(nums1) {
		a[k] = nums1[i]
		i += 1
		k += 1
	}

	for j < len(nums2) {
		a[k] = nums2[j]
		j += 1
		k += 1
	}

	r1 := float32(len(a)) / 2
	r2 := len(a) / 2

	var res float64
	if r1 > float32(r2) {
		res = float64(a[r2])
	} else {
		res = float64((a[r2-1] + a[r2])) / 2
	}

	return res
}
