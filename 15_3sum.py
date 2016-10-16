class Solution(object):
    def threeSum(self,nums):
        num = sorted(nums)
        x,y,ans = 0,1,[]
        while x < y:
            y = x + 1
            z = len(num) - 1
            while y < z:
                s = num[x] + num[y] + num[z]
                if s > 0:
                    z -= 1
                    while y < z and num[z] == num[z-1]:
                        z -= 1
                else:
                    if s == 0:
                        ans.append([num[x],num[y],num[z]])
                    y += 1
                    while y < z and num[y] == num[y-1]:
                        y += 1
            x += 1
            while x < len(num) and num[x] == num[x-1]:
                x += 1
        return ans
