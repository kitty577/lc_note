package main

// 回溯
func solveNQueens(n int) [][]string {
    // res 最终结果
    res:=[][]string{}
    
    // chessboard 记录矩阵每个位置放置的状态
    chessboard:=make([][]string,n)
    for i:=0;i<n;i++{
        chessboard[i]=make([]string,n)
    }
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            chessboard[i][j]="."
        }
    }
    var backtracking func(row int)
    backtracking=func(row int){
        // 已经全部放置完 退出
        if row==n{
            tmp:=make([]string,n)
            for i,rowStr:=range chessboard{
                tmp[i]=strings.Join(rowStr,"")
            }
            res=append(res,tmp)
            return
        }

        // 对每层的节点进行处理
        for i:=0;i<n;i++{
            if isValid(n,row,i,chessboard){
                chessboard[row][i]="Q"
                backtracking(row+1)
                chessboard[row][i]="."
            }
        }
    }

    backtracking(0)
    return res
}

func isValid(n,row,col int,chessboard [][]string)bool{
    // 列
    for i:=0;i<row;i++{
        if chessboard[i][col]=="Q"{
            return false
        }
    }

    // 斜线 \方向
    for i,j:=row-1,col-1;i>=0&&j>=0;i,j=i-1,j-1{
        if chessboard[i][j]=="Q"{
            return false
        }
    }

    // 斜线 /方向
    for i,j:=row-1,col+1;i>=0&&j<n;i,j=i-1,j+1{
        if chessboard[i][j]=="Q"{
            return false
        }
    }
    return true
}