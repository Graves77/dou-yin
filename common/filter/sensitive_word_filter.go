package filter

const replaceString = "***"

// 1.  定义Trie树节点：
type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// 2.  构造Trie树：
func buildTrieTree(words []string) *TrieNode {
	root := NewTrieNode()
	for _, word := range words {
		node := root
		for _, r := range word {
			if _, ok := node.children[r]; !ok {
				node.children[r] = NewTrieNode()
			}
			node = node.children[r]
		}
		node.end = true
	}
	return root
}

// 3.  对文本进行过滤：
func filterText(root *TrieNode, text string) string {
	var result []rune
	for i, r := range text {
		node := root
		j := i
		for node != nil {
			node = node.children[r]
			if node != nil && node.end {
				for ; j < len(text); j++ {
					result = append(result, '*')
				}
				break
			}
			if j < len(text)-1 {
				result = append(result, r)
				j++
				r = []rune(text)[j]
			} else {
				result = append(result, r)
				break
			}
		}
	}
	return string(result)
}
