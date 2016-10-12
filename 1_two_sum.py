#hash
class Solution(object):
    def twoSum(self, nums, target):
        hash = {} #dict()
        for x in range(len(nums)):
            if hash.get(targe - nums[x], None) == None:
                hash[nums[x]] == x
            else:
                return [hash[target - nums[x]] + 1 , x + 1]
