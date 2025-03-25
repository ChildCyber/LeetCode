package leetcode

// 接雨水
// https://leetcode-cn.com/problems/trapping-rain-water/

// 问题本质：
// 每个位置能接的雨水量由左右两边的最高柱子决定
// 水量 = min(左边最高, 右边最高) - 当前高度

// 找凹槽
// 时间复杂度：O(n²)
func trapFindGroove(height []int) int {
	// 思路：找谷底 → 配对两边墙
	// 1. 找出所有的“谷地”：左边有个高墙，右边也有个高墙，中间是凹进去的地方
	// 2. 对于每个凹槽：
	//    左墙高度 = 遇到的最高柱子（往左找）
	//    右墙高度 = 遇到的最高柱子（往右找）
	//    凹槽能装的水高度 = min(左墙, 右墙) - 谷底
	//    宽度 = 两墙之间的距离
	//    累加出水量
	// 3. 最后把所有小凹槽的水量加起来
	n := len(height)
	if n < 3 {
		return 0
	}

	res := 0
	for i := 1; i < n-1; i++ { // 第一个和最后一个元素只能作为边界的墙
		// 找到谷底：height[i] 必须小于左右
		if height[i] < height[i-1] && height[i] < height[i+1] { // 和暴力法的唯一区别
			// 严格谷底条件太死:
			// 会漏的情况：
			//   “谷底”左边或右边是平的：比如 [3,1,1,3]，两个 1 都能积水，但 height[i] < height[i-1] 在第二个 1 不成立。
			//   不是严格谷底的凹槽：比如 [4,2,3]，中间的 2 能存水
			// 找左墙
			maxLeft := 0
			for l := 0; l <= i; l++ { // 把谷底包括在内不会出错，反而能避免处理边界情况
				if height[l] > maxLeft {
					maxLeft = height[l]
				}
			}

			// 找右墙
			maxRight := 0
			for r := i; r < n; r++ { // 把谷底包括在内不会出错，反而能避免处理边界情况
				if height[r] > maxRight {
					maxRight = height[r]
				}
			}

			// 算水
			h := min(maxLeft, maxRight) - height[i]
			if h > 0 {
				res += h
			}
		}
	}
	return res
}

// 暴力
// 时间复杂度：O(n²)
func trapBrute(height []int) int {
	// 思路：对于每个位置能存多少水，取决于它左边最高的墙和右边最高的墙中较矮的那个
	// 1. 遍历所有位置 i（忽略左右两端，因为边界接不了水）
	// 2. 对于每个 i，往左扫一遍，找到 maxLeft（左侧最高柱子）
	// 3. 再往右扫一遍，找到 maxRight（右侧最高柱子）
	// 4. 当前位置能接的水量 = min(maxLeft, maxRight) - height[i]（如果是负数，就说明比柱子还矮，不接水）
	// 5. 把所有位置的水量累加
	// 和找谷底的区别在于：暴力法把每个柱子都当成潜在“蓄水点”，无论它是不是谷底，都检查一遍
	n := len(height)
	if n == 0 {
		return 0
	}

	res := 0
	for i := 1; i < n-1; i++ { // 两端不用算
		maxLeft, maxRight := 0, 0

		// 往左找最高
		for l := 0; l <= i; l++ {
			if height[l] > maxLeft {
				maxLeft = height[l]
			}
		}

		// 往右找最高
		for r := i; r < n; r++ {
			if height[r] > maxRight {
				maxRight = height[r]
			}
		}

		// 当前能接的水
		water := min(maxLeft, maxRight) - height[i]
		if water > 0 {
			res += water
		}
	}
	return res
}

// 暴力-动态规划（预处理）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func trapPreprocess(height []int) int {
	// 暴力每次位置都要往两边扫， 预处理版把“左最高、右最高”提前存起来，查表即可
	n := len(height)
	if n == 0 {
		return 0
	}

	// 1. 预处理 maxLeft：maxLeft[i] 表示 [0..i] 范围内的最高柱子
	maxLeft := make([]int, n)
	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		if height[i] > maxLeft[i-1] {
			maxLeft[i] = height[i]
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}

	// 2. 预处理 maxRight：maxRight[i] 表示 [i..n-1] 范围内的最高柱子
	maxRight := make([]int, n)
	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		if height[i] > maxRight[i+1] {
			maxRight[i] = height[i]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	// 3. 计算水量
	res := 0
	for i := 1; i < n-1; i++ {
		water := min(maxLeft[i], maxRight[i]) - height[i]
		if water > 0 {
			res += water
		}
	}

	return res
}

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func trap(height []int) int {
	// 维护两个指针 left 和 right，以及两个变量（当前扫描过的最大高度） leftMax 和 rightMax
	// 初始时 left=0,right=n−1,leftMax=0,rightMax=0。
	// 指针 left 只会向右移动，指针 right 只会向左移动，在移动指针的过程中维护两个变量 leftMax 和 rightMax 的值。
	// 当两个指针没有相遇时，进行如下操作：
	// 使用 height[left] 和 height[right] 的值更新 leftMax 和 rightMax 的值；
	// 如果 height[left]<height[right]，则必有 leftMax<rightMax，下标 left 处能接的雨水量等于 leftMax−height[left]，将下标 left 处能接的雨水量加到能接的雨水总量，然后将 left 加 1（即向右移动一位）；
	// 如果 height[left]≥height[right]，则必有 leftMax≥rightMax，下标 right 处能接的雨水量等于 rightMax−height[right]，将下标 right 处能接的雨水量加到能接的雨水总量，然后将 right 减 1（即向左移动一位）。
	// 当两个指针相遇时，即可得到能接的雨水总量。
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right { // 两指针相遇
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++ // 向右移动一位
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right-- // 向左移动一位
		}
	}
	return res
}

// 单调栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func trapStack(height []int) int {
	// 单调栈可以用来维护“递减高度的下标”，一旦出现更高的柱子，就触发结算：低洼区域被填满水了
	n := len(height)
	if n == 0 {
		return 0
	}

	stack := []int{} // 存下标。栈里的柱子高度是单调递减的
	res := 0

	for i := 0; i < n; i++ {
		// 当前高度比栈顶大，触发结算
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1] // 低谷位置
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break // 左边没墙，不能装水
			}

			left := stack[len(stack)-1]
			h := min(height[left], height[i]) - height[mid]
			w := i - left - 1
			if h > 0 {
				res += h * w
			}
		}
		// 如果当前高度小于等于栈顶 ，继续入栈（说明还在往低谷走）
		stack = append(stack, i)
	}

	return res
}
