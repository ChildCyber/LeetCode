package leetcode

import "sort"

// 多数元素
// https://leetcode-cn.com/problems/majority-element/

// 哈希表
// 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement(nums []int) int {
	// 存储每个元素以及出现的次数。键表示一个元素，值表示该元素出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 { // 遍历数组时直接判断，省去了对哈希表的遍历
			return v
		}
	}
	return 0
}

// 排序法
// 将数组排序，中间位置的数字一定是多数元素（因为多数元素超过一半，所以排序后肯定在中间）
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票法(Boyer-Moore)
// 就像是多轮淘汰赛：不同的数字相互"抵消"，最后剩下的就是多数元素
// 时间复杂度：O(n)
// 空间复杂度: O(1)
func majorityElementMoore(nums []int) int {
	candidate, count := nums[0], 0
	// 遍历元素，对于数组中的每个元素：
	for i := 0; i < len(nums); i++ {
		if count == 0 { // 如果计数器为0，将当前元素设为候选元素
			candidate, count = nums[i], 1
		} else if nums[i] == candidate { // 如果当前元素等于候选元素，计数器加1
			count++
		} else { // 如果当前元素不等于候选元素，计数器减1
			count--
		}
	}
	// 遍历结束后，候选元素就是可能的多数元素
	return candidate
}
