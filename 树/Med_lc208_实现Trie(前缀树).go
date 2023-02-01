package main

// 每个节点 都存储着它的子节点信息
type Trie struct {
    children [26]*Trie
    isEnd bool   // 表示当前节点是否是一个单词的结束
}

func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    node:=this
    for _,ch:=range word{
        ch=ch-'a'
        if node.children[ch]==nil{   // 不存在，开辟新节点
            node.children[ch]=&Trie{}
        }
        node=node.children[ch]
    }
    node.isEnd=true
}

func (this *Trie)SeerchPrefix(prefix string)*Trie{
    node:=this
    for _,ch:=range prefix{
        ch=ch-'a'
        if node.children[ch]==nil{
            return nil
        }
        node=node.children[ch]
    }
    return node
}

func (this *Trie) Search(word string) bool {
    node:=this.SeerchPrefix(word)
    return node!=nil&&node.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    return this.SeerchPrefix(prefix)!=nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */


// 2023.01.03