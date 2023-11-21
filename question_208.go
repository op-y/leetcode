package leetcode

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	nchar    byte
	isWord   bool
	parent   *TrieNode
	children []*TrieNode
}

func NewTrieNode(c byte, p *TrieNode) *TrieNode {
	children := make([]*TrieNode, 26)
	return &TrieNode{
		nchar:    c,
		parent:   p,
		children: children,
	}
}

func Constructor208() Trie {
	root := NewTrieNode('0', nil)
	return Trie{root: root}
}

func (this *Trie) Insert(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		if p.children[word[i]-'a'] == nil {
			// 当前这个字符在路径上不存在
			child := NewTrieNode(word[i], p)
			p.children[word[i]-'a'] = child
			p = child
		} else {
			p = p.children[word[i]-'a']
		}
	}
	p.isWord = true
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for i := 0; i < len(word); i++ {
		p = p.children[word[i]-'a']
		if p == nil {
			return false
		}
	}
	return p.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for i := 0; i < len(prefix); i++ {
		p = p.children[prefix[i]-'a']
		if p == nil {
			return false
		}
	}
	return true
}
