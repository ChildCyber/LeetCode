package leetcode

// 找到所有数组中消失的数字
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/

// 鸽笼原理
func findDisappearedNumbers(nums []int) []int {
	ans := []int{}
	for _, v := range nums { // 1～n的位置表示1～n个笼子，如果出现过，相应的鸽笼就会被占掉，将数值置为负数表示被占
		if v < 0 { // 当前遍历到的元素作为下标，下标需要为正数
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1] // nums[i]在区间[1, n]内，下标范围在[0,n-1]内，需要减1
		}
	}
	for i, v := range nums { // 遍历，正数就是没出现的数字
		if v > 0 {
			ans = append(ans, i+1) // 缺失的数为下标加1
		}
	}
	return ans
}
