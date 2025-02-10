func removeElement(nums []int, val int) int {
	k := 0
	
	for _, curr := range nums {
		if curr != val {
			nums[k] = curr
			k++
		}
	}

	return k
}