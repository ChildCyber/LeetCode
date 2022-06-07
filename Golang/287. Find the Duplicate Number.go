package leetcode

// 寻找重复数
// https://leetcode-cn.com/problems/find-the-duplicate-number/

// set
// 不满足空间复杂度O(1)

// 快慢指针
// 将数字想象成链表中的结点，数组中数字代表下一个结点的数组下标。找重复的数字就是找链表中成环的那个点。
// 对 nums 数组建图，每个位置 i 连一条 i→nums[i] 的边。由于存在的重复的数字 target，因此 target 这个位置一定有起码两条指向它的边，因此整张图一定存在环，且要找到的 target 就是这个环的入口
// Floyd 判圈算法，检测链表是否有环
func findDuplicate(nums []int) int {
	// 先设置慢指针 slow 和快指针 fast，慢指针每次走一步，快指针每次走两步，两个指针在有环的情况下一定会相遇，此时再将 slow 放置起点 0，两个指针每次同时移动一步，相遇的点就是答案。
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow { // 快慢指针相遇
		slow = nums[slow]       // 慢指针一次走 1 步
		fast = nums[nums[fast]] // 快指针一次走 2 步
	}
	// 相交后，快指针从头开始，每次走一步，再次遇⻅的时候就是成环的交点处，也即是重复数字所在的地方
	walker := 0
	for walker != slow { // 重新相遇
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

// 二分查找
func findDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
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
