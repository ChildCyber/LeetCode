package leetcode

// K和数对的最大数目
// https://leetcode.cn/problems/max-number-of-k-sum-pairs/

// 哈希表
func maxOperations(nums []int, k int) int {
	count := 0
	m := make(map[int]int)

	for _, num := range nums {
		target := k - num
		if m[target] > 0 {
			// 找到一对
			count++
			m[target]--
		} else {
			// 没有配对，存起来
			m[num]++
		}
	}

	return count
}
