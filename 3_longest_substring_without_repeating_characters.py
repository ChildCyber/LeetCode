# 维护两个指针，保证前后两个指针之间的串没有重复字符，后指针扫到某个字符重复时就将前指针向后移到第一个和当前字符相同的字符之后
class Solution(object):
    def lengthOfLongestSubstring(self,s):
        hash,ans,pointer1,pointer2 = {},0,0,0
        while pointer2 < len(s):
            pointer = hash.get(s[pointer2],None)
            if pointer == None:
                hash[s[pointer2]] = pointer2
                pointer2 += 1
                ans = max(ans,pointer2 - pointer1)
            else:
                while pointer1 <= pointer:
                    hash.pop(s[pointer1])
                    pointer1 += 1
                pointer1 = pointer + 1
        return ans
