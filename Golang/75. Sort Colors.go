package leetcode

// 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	r, w, b := 0, 0, 0
	for _, num := range nums {
		if num == 0 { // 使用整数 0、 1 和 2 分别表示红色、白色和蓝色
			nums[b] = 2
			b++
			nums[w] = 1
			w++
			nums[r] = 0
			r++
		} else if num == 1 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
		} else if num == 2 {
			b++
		}
	}
}
