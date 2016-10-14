#回文数
class Solution(object):
    def isPalindrome(self,n):
        if n < 0:
            return False
        return str(n) == str(n)[::-1]
        
class Solution(object):
    def isPalindrome(self,n):
        if n <= 0:
            return False if n < 0 else True
        x,y = n,0
        while x:
            y,x = y * 10 + x % 10,x / 10
        return y == n
