#参考https://discuss.leetcode.com/topic/53883/short-python-solution
#pythonic
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

class Solution(object):
    def longestCommonPrefix(self, strs):
        if not strs:
            return ''
        pos = 0
        random = strs[0]
        for x in range(len(random)):
            for y in strs:
                if x >= len(y) or random[x] != y[x]:
                    pos = x
                    return random[:pos]
        return random
