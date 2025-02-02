from typing import List


class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        checked = set(nums)

        return len(nums) != len(checked)
