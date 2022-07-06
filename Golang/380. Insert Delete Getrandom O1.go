package leetcode

import "math/rand"

// O(1) 时间插入、删除和获取随机元素
// https://leetcode.cn/problems/insert-delete-getrandom-o1/
// 变长数组 + 哈希表
type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	// 判断 val 是否在哈希表中，如果已经存在则返回 false，如果不存在则插入 val
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums) // 变长数组长度为 val 的下标 index，将 val 和下标 index 存入哈希表
	rs.nums = append(rs.nums, val) // 在变长数组的末尾添加val
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	// 判断 val 是否在哈希表中，如果不存在则返回 false，如果存在则删除 val
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	// 从哈希表中获得 val 的下标 index
	last := len(rs.nums) - 1
	// 将变长数组的最后一个元素 last 移动到下标 index 处，在哈希表中将 last 的下标更新为 index
	rs.nums[id] = rs.nums[last]
	// 在变长数组中删除最后一个元素，在哈希表中删除 val
	rs.indices[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.indices, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
