package leetcode

// 实现 Trie (前缀树)
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// 应用场景：自动补全、拼写检查

// 一棵有根树，其每个节点包含以下字段：
//  指向子节点的指针数组children。对于本题而言，数组长度为26，即小写英文字母的数量。此时children[0]对应小写字母a，…，children[25]应小写字母z。
//  布尔字段 isWord，表示该节点是否为字符串的结尾。
type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

func Constructor208() Trie {
	return Trie{
		isWord:   false,
		children: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	// 从字典树的根开始，插入字符串。对于当前字符对应的子节点，有两种情况：
	// 1. 子节点存在。沿着指针移动到子节点，继续处理下一个字符。
	// 2. 子节点不存在。创建一个新的子节点，记录在 children数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。
	// 重复以上步骤，直到处理字符串的最后一个字符，然后将当前节点标记为字符串的结尾。
	parent := this
	for _, ch := range word { // 遍历插入字符串
		if child, ok := parent.children[ch]; ok { // 子节点存在
			parent = child
		} else { // 子节点不存在
			newChild := &Trie{
				children: make(map[rune]*Trie),
			}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func (this *Trie) Search(word string) bool {
	// 从字典树的根开始，查找前缀。对于当前字符对应的子节点，有两种情况：
	// 1. 子节点存在。沿着指针移动到子节点，继续搜索下一个字符。
	// 2. 子节点不存在。说明字典树中不包含该前缀，返回空指针。
	// 重复以上步骤，直到返回空指针或搜索完前缀的最后一个字符。
	parent := this
	for _, ch := range word { // 遍历插入字符串
		if child, ok := parent.children[ch]; ok { // 子节点存在
			parent = child
			continue
		}
		return false // 子节点不存在
	}
	return parent.isWord // 完全匹配word，当前节点是否为word的最后一个字符
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}
