package staticarrays

func removeDuplicates(nums []int) int {
	n := len(nums)

	left, right := 0, 0

	for right < n {
		nums[left] = nums[right]
		for right < n && nums[left] == nums[right] {
			right++
		}
		left++
	}

	return left
}

func removeDuplicates2(nums []int) int {
	left := 1

	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}
