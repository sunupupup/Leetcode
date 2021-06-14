package main


//我的方法  时间100%  内存 6%
func spiralOrder(matrix [][]int) (ret []int) {
    top,left := 0,0
    bottom,right := len(matrix)-1,len(matrix[0])-1

    //ret = make([]int,0,(bottom+1)*(right+1))    

    for top<=bottom&&left<=right {

        //1. top层   从左到右
        index := left
        for index<=right {
            ret = append(ret,matrix[top][index])
            index++
        }
        top++
        if !isContinue(top,bottom,left,right){
            break
        }
        //2.right层  从上到下
        index = top
        for index<=bottom {
            ret = append(ret,matrix[index][right])
            index++
        }
        right--
        if !isContinue(top,bottom,left,right){
            break
        }
        //3.bottom层 从右到左
        index = right
        for index>=left {
            ret = append(ret,matrix[bottom][index])
            index--
        }
        bottom--
        if !isContinue(top,bottom,left,right){
            break
        }
        //4. left层  从下到上
        index = bottom
        for index>=top {
            ret = append(ret,matrix[index][left])
            index--
        }
        left++
    }
    return ret
}

func isContinue(top,bottom,left,right int) bool {
    return top<=bottom&&left<=right 
}

//官方的层层输出
func spiralOrder2(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    var (
        rows, columns = len(matrix), len(matrix[0])
        order = make([]int, rows * columns)
        index = 0
        left, right, top, bottom = 0, columns - 1, 0, rows - 1
    )

    for left <= right && top <= bottom {
        for column := left; column <= right; column++ {
            order[index] = matrix[top][column]
            index++
        }
        for row := top + 1; row <= bottom; row++ {
            order[index] = matrix[row][right]
            index++
        }
        if left < right && top < bottom {
            for column := right - 1; column > left; column-- {
                order[index] = matrix[bottom][column]
                index++
            }
            for row := bottom; row > top; row-- {
                order[index] = matrix[row][left]
                index++
            }
        }
        left++
        right--
        top++
        bottom--
    }
    return order
}
