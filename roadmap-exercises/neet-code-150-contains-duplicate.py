from typing import List


class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        numsMap = {}
        for i in range(len(nums)):
            if nums[i] in numsMap:
                return True
            else:
                numsMap[nums[i]] = True

        return False
