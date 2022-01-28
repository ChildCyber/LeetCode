# -*- coding: utf-8 -*-
# python 3
# 需要考虑溢出

# Description
#
# Reverse digits of an integer.


class Solution(object):
    def reverse(self, n):
        """
        :type n: int
        :rtype: int
        """
        x = 0
        y = n if n > 0 else -n
        while y:
            x, y = x * 10 + y % 10, y // 10
        if x > 2 ** 31 or x < -(2 ** 31):
            return 0
        return x if n > 0 else -x


if __name__ == '__main__':
    print Solution().reverse(123)
    print Solution().reverse(-123)

