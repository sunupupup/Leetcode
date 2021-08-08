/*
1137. 第 N 个泰波那契数
泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。


*/

//复杂度O（n）
func tribonacci(n int) int {
	first := 0
	second := 1
	third := 1
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	for i := 0; i <= n-3; i++ {
		first, second, third = second, third, first+second+third
	}

	return third

}

//复杂度O（ logn ）
//矩阵快速幂
type matrix [3][3]int

func (a matrix) mul(b matrix) matrix {
    c := matrix{}
    for i, row := range a {
        for j := range b[0] {
            for k, v := range row {
                c[i][j] += v * b[k][j]
            }
        }
    }
    return c
}

func (a matrix) pow(n int) matrix {
    res := matrix{}
    for i := range res {
        res[i][i] = 1
    }
    for ; n > 0; n >>= 1 {
        if n&1 > 0 {
            res = res.mul(a)
        }
        a = a.mul(a)
    }
    return res
}

func tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n <= 2 {
        return 1
    }
    m := matrix{
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
    }
    res := m.pow(n)
    return res[0][2]
}
