package main

// 由于是字符串构造前缀树，所以可以用[26]int 记录字符串
type Trie struct {
	children [26]*Trie
	isEnd    bool // 到此节点上，是否是一个单词的结束
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch = ch - 'a'
		if node.children[ch] == nil { // 不存在此字符的结点，则为此开辟新节点
			node.children[ch] = &Trie{
				children: [26]*Trie{},
				isEnd:    false,
			}
		}
		node = node.children[ch]

	}
	node.isEnd = true // 单词遍历结束，标记
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		ch = ch - 'a'
		if node.children[ch] == nil {
			return false
		} else {
			node = node.children[ch]
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		ch = ch - 'a'
		if node.children[ch] == nil {
			return false
		} else {
			node = node.children[ch]
		}
	}
	return true
}
