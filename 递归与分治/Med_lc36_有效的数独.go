package main

// 二维数组记录行、列中数字出现的情况
// 三维数组记录3*3宫内数字出现的情况
func isValidSudoku(board [][]byte) bool {
    var rows,cols [9][9]int
    var group [3][3][9]int

    for i,row:=range board{
        for j,v:=range row{
            if v=='.'{
                continue
            }

            index:=v-'1'
            rows[i][index]++
            cols[j][index]++
            group[i/3][j/3][index]++
            if rows[i][index]>1||cols[j][index]>1||group[i/3][j/3][index]>1{
                return false
            }
        }
    }
    return true
}