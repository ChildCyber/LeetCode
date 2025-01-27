package leetcode

// 跳跃游戏
// https://leetcode-cn.com/problems/jump-game/
// 贪心
func canJump(nums []int) bool {
	if len(nums) == 0 { // 没有起点，也没有终点
		return false
	}
	if len(nums) == 1 { // 起点就是终点，不需要任何跳跃
		return true
	}

	maxJump := 0 // 当前能跳到最远的距离（索引）
	for i, v := range nums {
		if i > maxJump { // 如果中间有一个点比 maxJump 还要大，说明在这个点和 maxJump 中间连不上了，有些点不能到达最后一个位置
			return false
		}
		maxJump = max(maxJump, i+v) // i+v代表当前位置i所能到达的最远距离
	}
	return true
}

func canJump1(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if nums[i]+i > maxReach {
			maxReach = nums[i] + i
		}
	}
	return maxReach >= len(nums)-1
}
