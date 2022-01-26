package leetcode

import "sort"

// 根据字符出现频率排序
// https://leetcode-cn.com/problems/sort-characters-by-frequency/
// 先统计每个字符的频次,然后排序,最后按照频次从高到低进行输出
func freqencySort(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}   // 字符频次，字符：频次
	cMap := map[int][]byte{} // 频次：字符列表
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys))) // 字符出现频次高到低排序
	res := make([]byte, 0)
	for _, k := range keys { // 该字符出现k次
		for i := 0; i < len(cMap[k]); i++ { // 出现k次的字符列表，迭代
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}
