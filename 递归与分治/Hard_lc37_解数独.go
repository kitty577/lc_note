package main

// 二维数组 记录行、列中数字出现的情况
// 三维数组记录3*3宫内数字出现的情况
// 另外需要一个数组，记录需要填充的空格坐标位置
func solveSudoku(board [][]byte)  {
    var rows,cols [9][9]bool
    var group [3][3][9]bool
    var space [][2]int

    for i,row:=range board{
        for j,v:=range row{
            if v=='.'{
                space=append(space,[2]int{i,j})
            }else{
                index:=v-'1'
                rows[i][index]=true
                cols[j][index]=true
                group[i/3][j/3][index]=true
            }
        }
    }

    var backtracking func(pos int)bool
    backtracking=func(pos int)bool{
        // 填充完了所有空格
        if pos==len(space){
            return true
        }
        // 开始处理空格，对每个空格进行填充0~9 回溯
        i,j:=space[pos][0],space[pos][1]
        for digit:=byte(0);digit<9;digit++{
            if !rows[i][digit]&&!cols[j][digit]&&!group[i/3][j/3][digit]{
                rows[i][digit]=true
                cols[j][digit]=true
                group[i/3][j/3][digit]=true
                board[i][j]=digit+'1'

                // 回溯
                if backtracking(pos+1){
                    return true
                }

                // 还原现场
                rows[i][digit]=false
                cols[j][digit]=false
                group[i/3][j/3][digit]=false
            }
        } 
        return false     
    }

    backtracking(0)
}