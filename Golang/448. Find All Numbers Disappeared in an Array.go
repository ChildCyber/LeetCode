package leetcode

// 找到所有数组中消失的数字
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, v := range nums { // 鸽笼原理，1～n的位置表示1～n个笼子，如果出现过，相应的鸽笼就会被占掉，将数值置为负数表示被占
		if v < 0 { // abs
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums { // 遍历，鸽笼为整数就是没出现的数字
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
