package leetcode

// 只出现一次的数字
// https://leetcode-cn.com/problems/single-number/

// 异或
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func singleNumber(nums []int) int {
	// 对数组中所有数字进行异或，最后剩下的数字就是只出现依次的数字
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

// 空间复杂度O(n)
// 1. 使用集合存储数字。遍历数组中的每个数字，如果集合中没有该数字，则将该数字加入集合；如果集合中已经有该数字，则将该数字从集合中删除，最后剩下的数字就是只出现一次的数字。
// 2. 使用哈希表存储每个数字和该数字出现的次数。遍历哈希表，得到只出现一次的数字。
// 3. 使用集合存储数组中出现的所有数字，并计算数组中的元素之和。用集合中的元素之和的两倍减去数组中的元素之和，该数字就是数组中只出现一次的数字。
