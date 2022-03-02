package leetcode

// hard
// 分发糖果
// https://leetcode-cn.com/problems/candy/
// 两次遍历
func candy(ratings []int) (ans int) {
	// 要求：
	// 1. 每个孩子至少分配到 1 个糖果
	// 2. 相邻两个孩子评分更高的孩子会获得更多的糖果

	// 将「相邻的孩子中，评分高的孩子必须获得更多的糖果」拆分为两个规则，分别处理
	// 左规则：当 ratings[i−1]<ratings[i] 时，i 号学生的糖果数量将比 i−1 号孩子的糖果数量多
	// 右规则：当 ratings[i]>ratings[i+1] 时，i 号学生的糖果数量将比 i+1 号孩子的糖果数量多
	// 遍历数组两次，处理出每一个学生分别满足左规则或右规则时，最少需要被分得的糖果数量。每个人最终分得的糖果数量即为这两个数量的最大值
	n := len(ratings)
	left := make([]int, n)
	// 第一次遍历
	// 左规则：当 ratings[i−1] < <ratings[i] 时，i 号学生的糖果数量将比 i−1 号孩子的糖果数量多
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	// 第二次遍历
	// 右规则：当 ratings[i]>ratings[i+1] 时，i 号学生的糖果数量将比 i+1 号孩子的糖果数量多
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

func candy1(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
