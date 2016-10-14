#python 3
#需要考虑溢出
class Solution(object):
    def reverse(self,x):
        isNegtive = False
        if x < 0:
            isNegtive = True
        x = int(str(x)[::-1])
        if x > 2 ** 31 or x < -(2 ** 31):
            return 0
        if isNegtive:
            return -x
        else:
            return x
            
class Solution(object):
    def reverse(self,n):
        x = 0
        y = n if n > 0 else -n
        while y:
            x, y = x * 10 + y % 10, y // 10
        if x > 2 ** 31 or x < -(2 ** 31):
            return 0
        return x if n > 0 else -x
