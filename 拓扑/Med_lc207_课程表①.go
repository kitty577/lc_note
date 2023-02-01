package main

func canFinish(numCourses int, prerequisites [][]int) bool {
    visited:=make([]int,numCourses)  // 记录结点搜索状态 0-未搜索，1-搜索中，2-搜索完成
    edges:= make([][]int, numCourses) // 存储有向图  前--》后
    valid:=true  // 合理
    var dfs func(int)

    for _,edge:=range prerequisites{
        edges[edge[1]]=append(edges[edge[1]],edge[0])
    }

    dfs=func(u int){
        visited[u]=1
        for _,v:=range edges[u]{
            if visited[v]==1{  // 形成环了
                valid=false
                return
            }else if visited[v]==0{
                dfs(v)
                if !valid{
                    return 
                }
            }
        }
        visited[u]=2  // 标记为搜索完成
    }

    for i:=0;i<numCourses;i++{
        if visited[i]==0{
            dfs(i)
        }
    }
    return valid

}

// 2023.01.11