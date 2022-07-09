package leetcode

// 四数相加 II
// https://leetcode.cn/problems/4sum-ii/
// 分组 + 哈希表
func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int, len(A)*len(B))
	// 对于 A 和 B，使用二重循环对它们进行遍历，得到所有 A[i]+B[j] 的值并存入哈希映射中。对于哈希映射中的每个键值对，每个键表示一种 A[i]+B[j]，对应的值为 A[i]+B[j] 出现的次数
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}

	ans := 0
	// 对于 C 和 D，使用二重循环对它们进行遍历。当遍历到 C[k]+D[l] 时，如果 −(C[k]+D[l]) 出现在哈希映射中，那么将 −(C[k]+D[l]) 对应的值累加进答案中
	for _, c := range C {
		for _, d := range D {
			ans += m[0-c-d]
		}
	}

	return ans
}
