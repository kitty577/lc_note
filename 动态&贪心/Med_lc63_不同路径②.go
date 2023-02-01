package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m:=len(obstacleGrid)
    n:=len(obstacleGrid[0])
    f:=make([][]int,m)
    for i:=0;i<m;i++{
        f[i]=make([]int,n)
    }
    for i:=0;i<m&&obstacleGrid[i][0]==0;i++{
        f[i][0]=1
    }
    for j:=0;j<n&&obstacleGrid[0][j]==0;j++{
        f[0][j]=1
    }

    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            if obstacleGrid[i][j]==1{   //障碍物
                continue
            }
            f[i][j]=f[i-1][j]+f[i][j-1]
        }
    }
    return f[m-1][n-1]
}

