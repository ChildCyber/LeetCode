# -*- coding: utf-8 -*-
# 回文数

# Description
#
# Determine whether an integer is a palindrome. Do this without extra space.


class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False
        return str(x) == str(x)[::-1]

    def isPalindrome2(self, n):
        if n <= 0:
            return False if n < 0 else True
        x, y = n, 0
        while x:
            y, x = y * 10 + x % 10, x / 10
        return y == n


if __name__ == '__main__':
    print Solution().isPalindrome(12321)
    print Solution().isPalindrome(12320)
    print Solution().isPalindrome(-12321)
    print Solution().isPalindrome2(12321)
    print Solution().isPalindrome2(12320)
    print Solution().isPalindrome2(-12321)

