package leetcode

// 煎饼排序
// https://leetcode-cn.com/problems/pancake-sorting/
// 每次排序都反转前 n 个数，n 小于数组的⻓度，最终数组是从小到大有序的

// 类选择排序
func pancakeSort(arr []int) (ans []int) {
	// 设一个元素的下标是 index，可以通过两次煎饼排序将它放到尾部：
	// 第一步选择 k=index+1，然后反转子数组 arr[0...k−1]，此时该元素已经被放到首部。
	// 第二步选择 k=n，其中 n 是数组 arr 的长度，然后反转整个数组，此时该元素已经被放到尾部。
	for n := len(arr); n > 1; n-- {
		index := 0
		for i, v := range arr[:n] {
			if v > arr[index] {
				index = i
			}
		}
		if index == n-1 {
			continue
		}
		// 第一次反转
		for i, m := 0, index; i < (m+1)/2; i++ {
			arr[i], arr[m-i] = arr[m-i], arr[i]
		}
		// 第二次反转
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		ans = append(ans, index+1, n)
	}
	return
}

// 递归
func pancakeSortRec(arr []int) (ans []int) {
	sort969(arr, len(arr), &ans)
	return
}

func sort969(arr []int, n int, ans *[]int) {
	if n == 1 {
		return
	}

	// 寻找最大饼的索引
	maxCake := 0
	maxCakeIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] > maxCake {
			maxCakeIndex = i
			maxCake = arr[i]
		}
	}
	// 第一次翻转，将最大饼翻到最上面
	reverse969(&arr, 0, maxCakeIndex)
	*ans = append(*ans, maxCakeIndex+1)
	// 第二次翻转，将最大饼翻到最下面
	reverse969(&arr, 0, n-1)
	*ans = append(*ans, n)
	// 递归调用
	sort969(arr, n-1, ans)
}

func reverse969(m *[]int, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
