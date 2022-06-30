package leetcode

import (
	"math"
	"sort"
)

// 最短无序连续子数组
// https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/
// 找出符 最短 子数组，并输出它的长度

// 排序
// 将给定的数组 nums 表示为三段子数组拼接的形式，分别记作 numsA​，numsB​，numsC
// 要求找到最短的 numsB，即最短逆序区间
// 将原数组 nums 排序与原数组进行比较，取最长的相同的前缀为 numsA​，取最长的相同的后缀为 numsC​，这样就可以取到最短的 numsB
func findUnsortedSubarraySort(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted) // 排序

	left, right := 0, len(nums)-1
	// 取最长的相同的前缀为 numsA​
	for nums[left] == numsSorted[left] {
		left++
	}
	// 取最长的相同的后缀为 numsC
	for nums[right] == numsSorted[right] {
		right--
	}

	return right - left + 1
}

// 一次遍历
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	// 把整个数组分成三段处理
	// 假设 numsB 在 nums 中对应区间为 [left,right]，不确保有序
	// numsA 在 nums 中对应区间为 [0:left]；numsC 在 nums 中对应区间为 [right:]，都满足升序要求
	left, right := -1, -1
	// 从大到小枚举 i，用一个变量 minn 记录，left 左侧即为 numsA 能取得的最大范围
	// 从小到大枚举 i，用一个变量 maxn 记录，right 右侧即为 numsC 能取得的最大范围
	minn, maxn := math.MaxInt64, math.MinInt64

	for i, num := range nums {
		// 如果当前值的前面还有更大的值，不满足升序，当前值应该被排序
		if maxn > num { // 从头到尾遍历数组，从小到大枚举i，当前值小于历史最大值，在进入right之前，遍历到的nums[i]都是小于maxn的
			right = i // right就是遍历中最后一个小于maxn元素的位置
		} else { // 当前最大值
			maxn = num
		}

		// 如果当前值的后面还有更小的值，不满足升序，当前值应该被排序
		if minn < nums[n-i-1] { // 从尾到头遍历数组，从大到大枚举i，当前值大于历史最小值，在进入left之前，那么遍历到的nums[i]也都是大于minn的
			left = n - i - 1 // left就是遍历中最后一个大于minn元素的位置
		} else { // 当前最小值
			minn = nums[n-i-1]
		}
	}

	// 特别判断nums有序，此时numsB的长度为0
	if right == -1 {
		return 0
	}

	return right - left + 1
}
