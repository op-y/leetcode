package leetcode

type WordDictionary struct {
	root *WDNode
}

type WDNode struct {
	nchar    byte
	isWord   bool
	parent   *WDNode
	children []*WDNode
}

func (n *WDNode) Search(word string) bool {
	if word == "" {
		return n.isWord
	}
	if word[0] == '.' {
		for _, child := range n.children {
			if child != nil {
				if ok := child.Search(word[1:]); ok {
					return true
				}
			}
		}
	} else {
		if n.children[word[0]-'a'] == nil {
			return false
		}
		return n.children[word[0]-'a'].Search(word[1:])
	}
	return false
}

func NewWDNode(c byte, p *WDNode) *WDNode {
	children := make([]*WDNode, 26)
	return &WDNode{
		nchar:    c,
		parent:   p,
		children: children,
	}
}

func Constructor211() WordDictionary {
	root := NewWDNode('0', nil)
	return WordDictionary{root: root}
}

func (this *WordDictionary) AddWord(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		if p.children[word[i]-'a'] == nil {
			// 当前这个字符在路径上不存在
			child := NewWDNode(word[i], p)
			p.children[word[i]-'a'] = child
			p = child
		} else {
			p = p.children[word[i]-'a']
		}
	}
	p.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.root.Search(word)
}
