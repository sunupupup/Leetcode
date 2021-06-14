package main

//二维数组的初始化
//     matrix := make([][]int, n)
//     for i := range matrix {
//	       matrix[i] = make([]int, n)
//     }

func generateMatrix(n int) [][]int {
    var ret [][]int
    ret = make([][]int,n)
    for i:=0;i<n;i++{
        ret[i] = make([]int,n)
    }

    top,left := 0,0
    bottom,right := n-1,n-1
    count := 1
    for top<bottom&&left<right{
        index := left
        for index<=right {
            ret[top][index] = count
            count++
            index++
        }
        top++

        index = top
        for index<=bottom {
            ret[index][right] = count
            count++
            index++
        }
        right--

        index = right
        for index>=left {
            ret[bottom][index] = count
            count++
            index--
        }
        bottom--
        
        index = bottom
        for index>=top {
            ret[index][left]=count
            count++
            index--
        }
        left++
    }

    if top==bottom {
        ret[top][top]=count
    }

    return ret
    
    
}