package main

// 情况：
// 支付5元 无需找零
// 支付10元 消耗一张5元 找零
// 支付20元 消耗一张10元和一张5元，或者 消耗三张5元
// ======》5元钞票很万能！找零先用完10元再用5元的
func lemonadeChange(bills []int) bool {
    // 记录各面值钞票的数量
    five,ten,twenty:=0,0,0

    for _,bill:=range bills{
        if bill==5{
            five++
        }
        if bill==10{
            if five<=0{
                return false
            }
            ten++
            five--
        }
        if bill==20{
            if ten>0&&five>0{
                ten--
                five--
                twenty++
            }else if five>=3{
                five=five-3
                twenty++
            }else{
                return false
            }
        }
    }

    return true
}