func twoSum(nums []int, target int) []int {
	for index, current := range nums {
        difference := target - current
        if indexDifference := indexOf(nums, difference); index != indexDifference && indexDifference >= 0 {
			return []int{index, indexDifference}
		}
	}
	return []int{0, 0}
}

func indexOf(arr []int, value int) int {
	for index, current := range arr {
		if current == value {
			return index
		}
	}
	return -1
}