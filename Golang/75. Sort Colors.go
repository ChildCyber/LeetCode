package leetcode

// 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
// 必须原地排序
// 计数排序
// 时间复杂度：O(2n) = O(n)
// 空间复杂度：O(1)
func sortColorsCount(nums []int) {
	// 思路：先统计每种元素出现的次数，然后按照顺序重新填充数组
	count0, count1, count2 := 0, 0, 0
	// 第一次遍历：统计 0、1、2 各自出现的次数
	for _, num := range nums {
		if num == 0 {
			count0++
		} else if num == 1 {
			count1++
		} else {
			count2++
		}
	}
	// 第二次遍历：按照 0→1→2 的顺序重新填充数组
	i := 0
	for j := 0; j < count0; j++ {
		nums[i] = 0
		i++
	}
	for j := 0; j < count1; j++ {
		nums[i] = 1
		i++
	}
	for j := 0; j < count2; j++ {
		nums[i] = 2
		i++
	}
}

// 三指针
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	// 使用三个指针将数组划分为三个区域：
	// left：指向已排序的 0 的末尾
	// curr：当前检查的元素
	// right：指向已排序的 2 的开头
	// 即[0,left)全为0，[left,curr)全为1，[curr,right)全为2

	// 算法步骤：
	// 1. 初始化 left = 0, curr = 0, right = n-1
	// 2. 当 curr <= right 时：
	//   如果 nums[curr] == 0：交换 nums[curr] 和 nums[left]，然后 left++, curr++
	//   如果 nums[curr] == 1：curr++（不做交换）
	//   如果 nums[curr] == 2：交换 nums[curr] 和 nums[right]，然后 right--
	left, curr, right := 0, 0, len(nums)-1

	for curr <= right {
		switch nums[curr] {
		case 0:
			// 交换到左边区域
			nums[curr], nums[left] = nums[left], nums[curr]
			left++
			curr++
		case 1:
			// 留在中间区域
			curr++
		case 2:
			// 交换到右边区域
			nums[curr], nums[right] = nums[right], nums[curr]
			right--
		}
	}
}
