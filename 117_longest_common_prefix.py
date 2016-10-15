#参考https://discuss.leetcode.com/topic/53883/short-python-solution
class Solution(object):
    def longestCommonPrefix(self, strs):
        if not strs:
            return ''
        shortest = min(strs,key=len)
        for x,y in enumerate(shortest):
            for z in strs:
                if z[x] != y:
                    return shortest[:x]
        return shortest
