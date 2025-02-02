class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        prevMap = {}

        for i, n in enumerate(nums):
            difference = target - n
            if difference in prevMap:
                return [prevMap[difference], i]
            prevMap[n] = i

        return []