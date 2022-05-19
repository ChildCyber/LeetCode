package leetcode

// 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
// 必须原地排序

// 三指针
// 统计出数组中 0,1,2 的个数，再根据它们的数量，重写整个数组
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	r, w, b := 0, 0, 0 // 红白蓝，分别指向最后一个0，1，2的下一个位置
	// [0,r)全为0，[r,w)全为1，[w,b)全为2
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
