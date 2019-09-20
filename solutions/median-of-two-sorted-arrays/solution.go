package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0, len(nums1)+len(nums2))

	var i, j int

	for i < len(nums1) && j < len(nums2) {
		var tmp int
		if nums1[i] > nums2[j] {
			tmp = nums2[j]
			j++
		} else {
			tmp = nums1[i]
			i++
		}
		nums = append(nums, tmp)
	}

	for ; i < len(nums1); i++ {
		nums = append(nums, nums1[i])
	}

	for ; j < len(nums2); j++ {
		nums = append(nums, nums2[j])
	}

	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / float64(2)
	} else {
		return float64(nums[len(nums)/2])
	}
}
