package main

// 能在常数时间内检索到最小元素
// 可以维护两个栈，一个栈正常操作，一个栈存储当前的最小值
type MinStack struct {
    normalStack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack{
        normalStack:[]int{},
        minStack:[]int{math.MaxInt64},
    }
}


func (this *MinStack) Push(val int)  {
    this.normalStack=append(this.normalStack,val)
    top:=this.minStack[len(this.minStack)-1]
    this.minStack=append(this.minStack,min(top,val))
    
}


func (this *MinStack) Pop()  {
    this.minStack=this.minStack[:len(this.minStack)-1]
    this.normalStack=this.normalStack[:len(this.normalStack)-1]
}


func (this *MinStack) Top() int {
    return this.normalStack[len(this.normalStack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}

func min(x,y int)int{
    if x>y{
        return y
    }
    return x
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */