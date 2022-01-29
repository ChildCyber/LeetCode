package leetcode

// 将有序数组转换为二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
// 中序遍历
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 递归
	return &TreeNode{
		Val:   nums[len(nums)/2], // 选取中间元素作为根节点，始终向下取整
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}
