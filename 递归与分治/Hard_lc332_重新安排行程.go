package main

// 可以使用map来记录映射关系（行程出发地和目的地）
// 同时要注意死循环问题、从一个出发地多次出发  ===>使用一个bool来标记是否已经飞过此段行程

type pair struct{
    target string
    visited bool
}

type pairs []*pair

func (ps pairs)Len()int{
    return len(ps)
}

func (ps pairs)Swap(i,j int){
    ps[i],ps[j]=ps[j],ps[i]
}

func (ps pairs)Less(i,j int)bool{
    return ps[i].target<ps[j].target
}

func findItinerary(tickets [][]string) []string {
    res:=[]string{}

    // map[出发机场] pairs{目的地，是否访问过}
    targets:=make(map[string]pairs)

    for _,ticket:=range tickets{
        if targets[ticket[0]]==nil{
            targets[ticket[0]]=make(pairs,0)
        }
        targets[ticket[0]]=append(targets[ticket[0]],&pair{ticket[1],false})
    }

    for k,_:=range targets{
        sort.Sort(targets[k])
    }

    res=append(res,"JFK")

    var backtracking func()bool
    backtracking=func()bool{
        if len(tickets)+1==len(res){
            return true
        }

        for _,pair:=range targets[res[len(res)-1]]{
            if pair.visited==false{
                res=append(res,pair.target)
                pair.visited=true

                if backtracking(){
                    return true
                }

                // 还原现场
                res=res[:len(res)-1]
                pair.visited=false
            }
        }
        return false
    }
    backtracking()

    return res
}