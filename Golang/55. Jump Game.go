package leetcode

// 跳跃游戏
// https://leetcode-cn.com/problems/jump-game/
// 贪心
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	maxJump := 0 // 能跳到最远的距离
	for i, v := range nums {
		if i > maxJump { // 如果中间有一个点比 maxJump 还要大，说明在这个点和 maxJump 中间连不上了，有些点不能到达最后一个位置
			return false
		}
		maxJump = max(maxJump, i+v) // i+v代表当前位置i所能到达的最远距离
	}
	return true
}
