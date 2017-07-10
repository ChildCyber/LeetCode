# hash

# Description
#
# Given an array of integers, return indices of the two numbers such that they add up to a specific target.
#
# You may assume that each input would have exactly one solution, and you may not use the same element twice.


class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        lookup = {}  # dict()
        for x in range(len(nums)):
            if lookup.get(target - nums[x], None) is None:
                lookup[nums[x]] = x
            else:
                return [lookup[target - nums[x]] + 1, x + 1]


if __name__ == '__main__':
    print Solution().twoSum((2, 7, 11, 15), 9)

