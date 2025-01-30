package leetcode

// 跳跃游戏 II
// https://leetcode-cn.com/problems/jump-game-ii/
// 贪心
func jump(nums []int) int {
	// 已经到达终点，不需要跳跃
	if len(nums) == 1 {
		return 0
	}
	// 跟踪三个信息：
	// 跳跃次数（jumps）：记录当前已经跳跃了多少次
	// 当前跳跃的边界（currentEnd）：表示当前跳跃所能到达的最远位置。遍历数组时，一旦到达这个边界，就意味着需要进行一次新的跳跃
	// 下一次跳跃的最远位置（farthest）：表示从当前点及其之前的点所能跳到的最远位置。每次遍历时，都需要更新这个值
	currentEnd, farthest, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		// 更新能够到达的最远位置
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// 每次到达当前跳跃的边界时，都会选择在边界之前的所有点中能跳到的最远点作为下一次跳跃的边界。这确保了每次跳跃尽可能远，从而最小化总跳跃次数
		// 如果已经到达当前跳跃的边界
		if i == currentEnd {
			// 进行下一次跳跃
			currentEnd = farthest
			jumps++
			// 如果已经能够到达终点，提前结束
			if currentEnd >= len(nums)-1 {
				return jumps
			}
		}
	}
	return jumps
}
