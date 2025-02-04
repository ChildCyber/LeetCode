package leetcode

// 实现 Trie (前缀树)
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// 应用场景：自动补全、拼写检查

// Trie 前缀树结构
type Trie struct {
	root *TrieNode
}

// Constructor208 初始化前缀树
func Constructor208() Trie {
	return Trie{root: &TrieNode{}}
}

// TrieNode 表示前缀树的一个节点
type TrieNode struct {
	children [26]*TrieNode // 只考虑小写字母 a-z
	isEnd    bool
}

// Insert 插入一个单词
func (t *Trie) Insert(word string) {
	// 从根节点开始，按字母逐个向下走，如果路径不存在就新建节点
	node := t.root
	// 遍历单词的每个字符
	for _, ch := range word {
		idx := ch - 'a'
		// 为每个字符创建对应的子节点（如果不存在）
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	// 在单词结尾标记
	node.isEnd = true
}

// Search 精确匹配一个单词
func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.isEnd
}

// StartsWith 判断是否存在以 prefix 开头的单词
func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

// findNode 内部工具：沿着字符串走，返回最后的节点或 nil
func (t *Trie) findNode(s string) *TrieNode {
	node := t.root
	for _, ch := range s {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}
