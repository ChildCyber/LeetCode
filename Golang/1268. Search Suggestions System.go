package leetcode

import (
	"sort"
	"strings"
)

// 搜索推荐系统
// https://leetcode.cn/problems/search-suggestions-system/

func suggestedProducts(products []string, searchWord string) [][]string {
	// 先对产品列表进行排序
	sort.Strings(products)

	result := make([][]string, 0, len(searchWord))
	prefix := ""

	// 遍历搜索词的每个字符
	for _, char := range searchWord {
		prefix += string(char)
		suggestions := make([]string, 0, 3)

		// 在排序后的产品列表中查找匹配前缀的产品
		for _, product := range products {
			if strings.HasPrefix(product, prefix) {
				suggestions = append(suggestions, product)
				if len(suggestions) == 3 {
					break
				}
			}
		}

		result = append(result, suggestions)
	}

	return result
}

// 二分
func suggestedProductsBS(products []string, searchWord string) [][]string {
	sort.Strings(products)
	result := make([][]string, 0, len(searchWord))

	left, right := 0, len(products)-1

	for i := 0; i < len(searchWord); i++ {
		// 更新左边界：找到第一个匹配当前前缀的产品
		for left <= right && (len(products[left]) <= i || products[left][i] != searchWord[i]) {
			left++
		}

		// 更新右边界：找到最后一个匹配当前前缀的产品
		for left <= right && (len(products[right]) <= i || products[right][i] != searchWord[i]) {
			right--
		}

		// 收集当前范围内的最多3个产品
		suggestions := make([]string, 0, 3)
		if left <= right {
			for j := left; j <= right && j < left+3; j++ {
				suggestions = append(suggestions, products[j])
			}
		}

		result = append(result, suggestions)
	}

	return result
}
