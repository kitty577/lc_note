package main

// 前缀树
type tire struct {
	children [26]*tire
	isEnd    bool
}

type MagicDictionary struct {
	*tire
}

func Constructor() MagicDictionary {
	return MagicDictionary{&tire{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := this.tire
		for _, ch := range word {
			ch = ch - 'a'
			if cur.children[ch] == nil {
				cur.children[ch] = &tire{}
			}
			cur = cur.children[ch]
		}
		cur.isEnd = true
	}
}

func dfs(node *tire, searchWord string, diff bool) bool {
	if len(searchWord) == 0 {
		return diff && node.isEnd
	}
	c := searchWord[0] - 'a'
	if node.children[c] != nil && dfs(node.children[c], searchWord[1:], diff) { // 匹配 继续下一个字符
		return true
	}

	// 如果还有替换机会，那么可以把当前的字符当做要替换的，继续搜索后面的字符
	if !diff {
		for i, child := range node.children {
			if i != int(c) && child != nil && dfs(child, searchWord[1:], true) {
				return true
			}
		}
	}
	return false
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this.tire, searchWord, false)
}
