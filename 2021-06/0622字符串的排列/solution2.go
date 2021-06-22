//参考官方，简单点的回溯，回溯过程中去重
func permutation(s string) []string {
    ret := make([]string,0)
    t := []byte(s)
    //对slice进行排序
    sort.Slice(t,func(i,j int)bool{return t[i] < t[j]})

    n:= len(t)
    perm := make([]byte,0,n)
    visit := make([]bool,n)
    var backtrack func(int) 
    backtrack= func(i int){
        if i==n {
            ret = append(ret,string(perm))
            return 
        }

        for j,b := range visit {

//如果这个点访问过了，或者这个点与前面的相同（前面没访问到，就不准访问后面的）
            if b || j>0 && !visit[j-1] && t[j-1]==t[j]{
                continue
            }
            //访问到这个节点
            visit[j] = true
            perm = append(perm,t[j])
            backtrack(i+1)
            perm = perm[:len(perm)-1]
            visit[j] = false
        }

    }

    backtrack(0)

    return ret
}