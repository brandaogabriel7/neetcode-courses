func containsDuplicate(nums []int) bool {
    non_duplicates := make(map[int]bool)

    for _, value := range nums {
        if non_duplicates[value] == true {
            return true
        }
        non_duplicates[value] = true
    }
    return false
}