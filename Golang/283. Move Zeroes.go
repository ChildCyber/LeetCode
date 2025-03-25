package leetcode

// 移动零
// https://leetcode-cn.com/problems/move-zeroes/

// 冒泡
func moveZeroesBubble(nums []int) {
	for i := 0; i < len(nums); i++ {
		swapped := false // 标记本轮是否发生了交换
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] == 0 && nums[j+1] != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true // 发生了交换
			}
		}
		// 如果本轮没有发生交换，说明所有零已经移动到了末尾，可以提前终止
		if !swapped {
			break
		}
	}
}

// 快慢指针-交换法
func moveZeroes(nums []int) {
	// 快指针：遍历整个数组，寻找非零元素
	// 慢指针：标记下一个非零元素应该放置的位置
	slow, fast, n := 0, 0, len(nums)
	// 快指针遍历整个数组
	for fast < n {
		// 如果快指针指向的元素不为零：
		if nums[fast] != 0 {
			// 将快指针指向的元素与慢指针指向的元素交换
			nums[slow], nums[fast] = nums[fast], nums[slow]
			// 慢指针向前移动一位
			slow++
		}
		// 如果快指针指向的元素为零：不做任何操作，继续移动快指针
		fast++
	}
}

// 移动再补零
func moveZeroesPadZero(nums []int) {
	j := 0 // 下一个非零元素应该放置的位置

	// 第一次遍历：将所有非零元素移到前面
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	// 第二次遍历：将剩余位置补零
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
