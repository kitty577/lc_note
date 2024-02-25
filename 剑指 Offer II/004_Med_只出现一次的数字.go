package main

// 1、使用哈希表
// func singleNumber(nums []int) int {
//     freq := map[int]int{}
//     for _, v := range nums {
//         freq[v]++
//     }
//     for num, occ := range freq {
//         if occ == 1 {
//             return num
//         }
//     }
//     return 0 
// }

// 2、位运算
// 将所有数的二进制的 各个位数上的数相加 得到结果1
// 由于除了某个元素只出现一次 其他都出现三次
// 那么分别对结果1中各个位数上的数对3取余，结果就是答案
// 如：  5--->  1 0 1
//       5--->  1 0 1
//       5--->  1 0 1
//       7--->  1 1 1

//  和====》    4 1 4

//  对三取余：  1  1 1    ==》7 则为答案

func singleNumber(nums []int) int {
    res:=int32(0)
    for i:=0;i<32;i++{
        // 计算所有数二进制的第I位的和
        sum:=int32(0)
        for _,v:=range nums{
            sum+=(int32(v)>>i)&0x1
        }
        // 拼接
        res|=(sum%3)<<i
    }
    return int(res)
}