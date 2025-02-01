package leetcode

// 寻找重复数
// https://leetcode-cn.com/problems/find-the-duplicate-number/

// 快慢指针（Floyd判圈算法）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func findDuplicate(nums []int) int {
	// 思路：将数组视为链表，利用指针移动来检测环
	// 对于数组 nums，建立这样的映射关系：
	//   当前索引 i 看作链表节点
	//   数组值 nums[i] 看作指向下一个节点的指针

	// 阶段1：检测环
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]       // 慢指针一次走 1 步
		fast = nums[nums[fast]] // 快指针一次走 2 步
	}

	// 阶段2：找到环的入口（重复数字）
	// 相交后，快指针从头开始，每次走一步，再次遇⻅的时候就是成环的交点处，也即是重复数字所在的地方
	fast = 0
	for slow != fast { // 重新相遇
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 二分查找
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func findDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		// 更新low或high
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 二进制
func findDuplicate2(nums []int) int {
	n := len(nums)
	ans := 0
	bitMax := 31
	for ((n - 1) >> bitMax) == 0 {
		bitMax--
	}
	for bit := 0; bit <= bitMax; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if (nums[i] & (1 << bit)) > 0 {
				x++
			}
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}
