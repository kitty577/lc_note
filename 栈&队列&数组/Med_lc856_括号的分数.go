package main

// 括号，想到之前做过的有效括号。借助栈
// (1)遇到左括号，入栈
// (2)遇到右括号，比较栈顶元素
//       if 左括号-->此对括号为最小匹配单位 ()形式，
//                 TODO: 弹出左括号，得分1入栈；
//       if 数值-->此括号中嵌套了其他平衡串 
//                 TODO: 弹出栈顶元素（即数值），暂存temp。再次比较栈顶元素
//                          if 左括号--> (A)形式，得分2*tmp 
//                          if 数值--> AB形式，得分tmp+栈顶数值 

func scoreOfParentheses(s string) int {
    var stack=[]int{0}
    for _,c:=range s{
        if c=='('{
            // 为了方便存数字和字符，用0代表左括号
            stack=append(stack,0)
        }else{
            top:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            stack[len(stack)-1]+=max(2*top,1)
        }
    }
    return stack[0]
}

func max(a, b int) int {
    if a > b{
        return a
    }
    return b
}

//cpp
class Solution {
	public:
		int scoreOfParentheses(string s) {
			stack<int> st;
			for(int i = 0; i < s.size(); i ++) {
				// 左括号入栈
				if(s[i] == '(') {
					st.push(0);
				} else {// 右括号分情况判断
				// 如果栈顶是左括号，拿出左括号放入1   ()
					if(st.top() == 0) {
						st.pop();
						st.push(1);
					} else {// 把栈顶的所有数字加起来
						int num = 0;
						while(st.top() > 0 && !st.empty()) {
							num += st.top();
							st.pop();
						}
						// 还有左括号执行(A)操作
						if(!st.empty()) {
							num *= 2;
							st.pop();
						}
						st.push(num);
					}
				}
			}
			int sum = 0;
			while(!st.empty()) {
				cout<<st.top()<<endl;
				sum += st.top();
				st.pop();
			}
			return sum;
		}
	};
	
	

// 2022.10.09