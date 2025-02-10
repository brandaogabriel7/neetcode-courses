func removeDuplicates (nums []int) int {
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
