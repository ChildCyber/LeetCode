# -*- coding: utf-8 -*-
# 参考https://discuss.leetcode.com/topic/53883/short-python-solution
# pythonic

# Description
#
# Write a function to find the longest common prefix string amongst an array of strings.


class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ''
        shortest = min(strs, key=len)
        for x, y in enumerate(shortest):
            for z in strs:
                if z[x] != y:
                    return shortest[:x]
        return shortest

    def longestCommonPrefix2(self, strs):
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

if __name__ == '__main__':
    print Solution().longestCommonPrefix(["hello", "heaven", "heavy"])
    print Solution().longestCommonPrefix2(["hello", "heaven", "heavy"])

