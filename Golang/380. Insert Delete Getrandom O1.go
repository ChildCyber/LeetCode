package leetcode

import "math/rand"

// O(1) 时间插入、删除和获取随机元素
// https://leetcode.cn/problems/insert-delete-getrandom-o1/

// 变长数组+哈希表
type RandomizedSet struct {
	nums     []int       // 存储实际元素，支持O(1)随机访问
	indexMap map[int]int // 存储元素值到数组索引的映射
}

func Constructor380() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	// 判断 val 是否在哈希表中，如果已经存在则返回 false，如果不存在则插入 val
	if _, ok := rs.indexMap[val]; ok {
		return false
	}

	// 插入到数组末尾
	rs.nums = append(rs.nums, val)
	// 记录元素索引
	rs.indexMap[val] = len(rs.nums) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	// 判断 val 是否在哈希表中，如果不存在则返回 false，如果存在则删除 val
	idx, ok := rs.indexMap[val]
	if !ok {
		return false
	}

	// 从哈希表中获得 val 的下标 index
	lastIdx := len(rs.nums) - 1
	lastVal := rs.nums[lastIdx]

	// 将当前元素与最后一个元素交换
	rs.nums[idx] = lastVal
	// 更新被交换元素的索引
	rs.indexMap[lastVal] = idx

	// 删除最后一个元素
	rs.nums = rs.nums[:lastIdx]
	// 删除哈希表中的映射
	delete(rs.indexMap, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
