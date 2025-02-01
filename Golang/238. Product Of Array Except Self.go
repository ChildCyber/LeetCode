package leetcode

// 除自身以外数组的乘积
// https://leetcode-cn.com/problems/product-of-array-except-self/
// 不使用除法，且时间复杂度为O(n)

// 暴力
func productExceptSelfForce(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)

	for i := 0; i < length; i++ {
		product := 1
		for j := 0; j < length; j++ {
			if i == j {
				continue
			}
			product *= nums[j]
		}
		ans[i] = product
	}
	return ans
}

// 前缀积+后缀积
// 思路：对于给定索引 i，使用它左边所有数字的乘积（称为前缀积）乘以右边所有数字的乘积（称为后缀积），然后将每个位置的前缀积和后缀积相乘就可以得到答案
// 避免暴力法中的重复计算：在暴力解法中为每个元素 i 都重新计算了一遍除了 i 之外所有元素的乘积
// 前缀乘积：对于每个位置 i，计算从数组开头到 i-1 位置所有元素的乘积
// prefix[i] = nums[0] * nums[1] * ... * nums[i-1]
// 对于任意位置 i，它左边的乘积就是 prefix[i]；当i=0时，i-1为-1，没有元素，所以用1代替
// 后缀乘积：对于每个位置 i，计算从 i+1 位置到数组末尾所有元素的乘积
// suffix[i] = nums[i+1] * nums[i+2] * ... * nums[n-1]
// 对于任意位置 i，它右边的乘积就是 suffix[i]；当i=n-1时，i+1为n，没有元素，所以用1代替
func productExceptSelf(nums []int) []int {
	length := len(nums)
	// 前、后缀积长度是length，因为需要的是 不包括当前元素 的前、后缀积
	prefix, suffix, answer := make([]int, length), make([]int, length), make([]int, length)

	// 计算前缀乘积
	prefix[0] = 1
	for i := 1; i < length; i++ { // 从前向后
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	// 计算后缀乘积
	suffix[length-1] = 1
	for i := length - 2; i >= 0; i-- { // 从后向前
		suffix[i] = nums[i+1] * suffix[i+1]
	}
	// 组合结果
	for i := 0; i < length; i++ {
		answer[i] = prefix[i] * suffix[i]
	}
	return answer
}

func productExceptSelf1(nums []int) []int {
	n := len(nums)

	// 创建结果数组，初始化为1
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

	// 计算前缀乘积（从左到右）
	// ans[i] 表示 nums[0] * nums[1] * ... * nums[i-1]
	prefix := 1
	for i := 0; i < n; i++ {
		ans[i] = prefix   // 将当前的前缀乘积赋给结果数组
		prefix *= nums[i] // 更新前缀乘积，包含当前元素
	}

	// 计算后缀乘积并直接与结果相乘（从右到左）
	// 此时 ans[i] 已经包含前缀乘积，再乘以后缀乘积
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix  // 将后缀乘积乘到结果上
		suffix *= nums[i] // 更新后缀乘积，包含当前元素
	}

	return ans
}
