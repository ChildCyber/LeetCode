package leetcode

// 将有序数组转换为二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

// 中序遍历-递归二分
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 递归
	return &TreeNode{
		Val:   nums[len(nums)/2], // 选取中间元素作为根节点，始终向下取整
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]), // nums[len(nums)/2]作为根节点，右子树需要从nums[len(nums)/2+1]开始
	}
}

// 递归二分-闭包
func sortedArrayToBSTClosure(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// 选择中间位置左边的元素作为根
		mid := left + (right-left)/2
		// 选择中间位置右边的元素
		// mid := left + (right-left+1)/2

		root := &TreeNode{Val: nums[mid]}
		root.Left = build(left, mid-1)
		root.Right = build(mid+1, right)

		return root
	}

	return build(0, len(nums)-1)
}
