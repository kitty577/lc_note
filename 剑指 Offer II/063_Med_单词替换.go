package main

import "strings"

// 单词替换。 替换成词根。
// 参考上一题的前缀树思路

func replaceWords(dictionary []string, sentence string) string {
	type tire map[rune]tire
	root := tire{}
	for _, s := range dictionary {
		cur := root
		for _, c := range s {
			if cur[c] == nil {
				cur[c] = tire{}
			}
			cur = cur[c]
		}
		cur['#'] = tire{} // 一个词根构建前缀树完成
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, c := range word {
			if cur['#'] != nil { // 匹配到字典前缀了，用前缀代替单词
				words[i] = word[:j]
				break
			}
			if cur == nil { // 没有前缀匹配，单词保持原样
				break
			}

			cur = cur[c] // 当前字母匹配到了，继续处理下一个字母
		}
	}

	return strings.Join(words, " ")
}
