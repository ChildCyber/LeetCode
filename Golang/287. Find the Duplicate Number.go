package leetcode

// 寻找重复数
// https://leetcode-cn.com/problems/find-the-duplicate-number/
// 快慢指针
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

// 二分查找
// 二进制
