package leetcode

// 种花问题
// https://leetcode-cn.com/problems/can-place-flowers/

// 贪心
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			// 检查左边是否空/不存在
			emptyLeft := (i == 0) || (flowerbed[i-1] == 0)
			// 检查右边是否空/不存在
			emptyRight := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

			if emptyLeft && emptyRight {
				flowerbed[i] = 1 // 种一朵花
				count++
				if count >= n {
					return true
				}
			}
		}
	}
	return count >= n
}
