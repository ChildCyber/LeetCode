package leetcode

// 拥有最多糖果的孩子
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxVal {
			maxVal = candies[i]
		}
	}

	ans := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxVal {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
