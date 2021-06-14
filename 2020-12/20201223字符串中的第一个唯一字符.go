package main

func firstUniqChar(s string) int {
    n := len(s)
    hash := [26]int{}
    for _,c:= range s{		//下标和值 两个接受
        hash[c-'a']++
    }
    for i:=0;i<n;i++ {
        if hash[s[i]-'a']==1{
            return i
        }
    }
    return -1
}