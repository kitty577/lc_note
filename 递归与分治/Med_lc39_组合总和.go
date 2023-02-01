package main

func combinationSum(candidates []int, target int) [][]int {
    res:=[][]int{}  // 组合集

    tmp:=[]int{}  // 单次组合

    var dfs func(target,idx int)

    dfs=func(target,idx int){
        if idx==len(candidates){
            return
        }

        if target==0{
            res=append(res,append([]int{},tmp...))
            return
        }

        // 不选
        dfs(target,idx+1)

        // 选
        if target-candidates[idx]>=0{
            tmp=append(tmp,candidates[idx])
            // 由于可以重复使用 所以下一次搜索还要从此元素开始
            dfs(target-candidates[idx],idx)

            // 还原现场
            tmp=tmp[:len(tmp)-1]
        }
    }
    dfs(target,0)
    return res
}

// 2023.01.07