package main

// 与前缀树相似，此题是找后缀，可以将单词反转构建前缀树

type tire struct {
	children [26]*tire
	depth    int
}

func revertStr(str string) string {
	b := []byte(str)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// 将单词插入到前缀树中
func (t *tire) Insert(str string) {
	cur := t
	d := 1
	for _, ch := range str {
		ch = ch - 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = &tire{
				children: [26]*tire{},
				depth:    d, // d是树节点的深度，其实也就是当前前缀的长度
			}
		}
		d++
		cur = cur.children[ch]
	}
}

// 前缀树。 走到所有的叶子节点
// 比如： em emit 这两个，em是在emit树的前半段，
// 那走到emit这个树的末端，取到长度，就是完成了em emit 这两个树的压缩合并
func dfs(t *tire, res *int) {
	isOver := true
	for _, node := range t.children {
		if node != nil {
			dfs(node, res)
			isOver = false
		}
	}
	if isOver {
		*res += t.depth + 1
	}
}

func minimumLengthEncoding(words []string) int {
	t := &tire{}
	for _, word := range words {
		t.Insert(revertStr(word))
	}

	num := 0
	dfs(t, &num)
	return num
}
