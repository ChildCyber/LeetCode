package leetcode

import "math"

// 下一个更大元素 III
// https://leetcode-cn.com/problems/next-greater-element-iii
func nextGreaterElement(n int) int {
	arr := []int{}
	for n != 0 { // 这样拆分，会导致数字倒过来。所以后面的逻辑都是倒着的。如果看不惯可以将数组reverse一下
		arr = append(arr, n%10)
		n /= 10
	}

	idx := -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			idx = i + 1
			break
		}
	}
	// 如果所有数字都是递减的，说明不存在一个比这个数字更大的值了，返回-1
	if idx == -1 {
		return -1
	}
	// 找到4右边大于它的最小数字进行交换。然后把这个位置后面的所有数字做一个增序排列
	// 6, 5, 4, 8, 7, 6, 5, 2, 1
	// 6, 5, 5, 8, 7, 6, 4, 2, 1
	// 6, 5, 5, 1, 2, 4, 6, 7, 8
	// 从4的位置起，开始查找
	t, tidx := math.MaxInt32, -1
	for i := idx; i >= 0; i-- {
		if arr[i] > arr[idx] && arr[i] < t {
			t = arr[i]
			tidx = i
		}
	}
	// 交换位置
	arr[tidx], arr[idx] = arr[idx], arr[tidx]

	//将后面的所有数字递增，并计算结果
	reverse(arr[0:idx])
	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res = res*10 + arr[i]
	}
	return res
}
