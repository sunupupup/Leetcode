/*
小美的机器人
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小美在数轴上放置了若干个机器人，这些机器人每到整数时刻，就会检查是否和其他机器人重合。如果重合，它就会原地爆炸。

这些机器人的移动速度均为 1 。举例来说，如果一个机器人初始位于点3，然后它的方向是向右的，则时刻1会位于点4，时刻2会位于点5。

小美现在给小团这样一个任务：小美将给出所有机器人的初始位置和初始朝向。小团的任务是判断每个机器人的爆炸时刻。当然，如果有一些机器人永远不会爆炸，则输出-1。

小团向你求助。你能帮帮小团吗？

注意事项1：一个机器人爆炸了之后，就不会再继续存在在这个数轴上。

举例来说，如果有三个机器人，一个位于位置0，向右，一个位于位置2，向右，一个位于位置4，向左。则时刻1的时候，后两个机器人会在位置3相遇并发生爆炸，此后第一个机器人和第三个机器人不会在时刻2继续爆炸（因为此时已经不存在第三个机器人了）

注意事项2：请注意，只有整数时刻机器人才会检查重合。

举例来说，如果有两个机器人，一个位于位置1，向右，一个位于位置2，向左，则它们并不会在整数时刻重合。因此它们两个不存在相遇爆炸。

注意事项3：保证机器人初始时刻不会重叠。换句话说，不存在在时刻0就立刻爆炸的机器人。



输入描述
第一行一个正整数 n 表示有 n 个机器人。

接下来 n 行，每行一个正整数和一个字符，以空格分隔。正整数代表机器人的坐标，字符为大写字母 L 和 R 的其中一个，分别表示机器人向左运动 和 向右运动。

输出描述
输出 n 行，每行一个数字，对应表示每个机器人的答案：

若该机器人会爆炸，输出爆炸时间；若该机器人不会爆炸，输出-1。
*/
package main

import (
	"fmt"
)

type robot struct {
	no         int    //机器人编号
	fangxiang  string //运动方向
	initIndex  int
	nowIndex   int
	bloomIndex int //爆炸时的位置
}

func main() {

	//先构建机器人数组，构建两个数组，一个往左，一个往右
	left := []robot{}
	right := []robot{}
	T := 0
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		index := 0
		s := ""
		fmt.Scan(&index)
		fmt.Scan(&s)
		if s == "L" {
			left = append(left, robot{t, s, index, index, -1})
		} else {
			right = append(right, robot{t, s, index, index, -1})
		}
	}

	//每次循环，增加检查是否有爆炸的，有的话就标记，并且检查接下来是否还会有爆炸现象

	//只要有一个不全炸了 , 两个里面必须都有-1存在，且right里面最小的位置  必须小于  left里面最大的位置
	for !checkIsAllBoom(&left) && !checkIsAllBoom(&right) && check2(&right, &left) {
		for i := range right {
			right[i].nowIndex++
		}
		for i := range left {
			left[i].nowIndex--
		}

		//检查没有爆炸的中，有没有重合爆炸的

		for i := range right {
			if right[i].bloomIndex != -1 {
				continue
			}
			index := right[i].nowIndex
			for j := range left {
				if left[j].bloomIndex == -1 && left[j].nowIndex == index {
					//这地方炸了
					left[j].bloomIndex = index
					right[i].bloomIndex = index
					break
				}
			}
		}

	}
	ret := make([]int, T)
	for i := range left {
		if left[i].bloomIndex == -1 {
			ret[left[i].no] = -1
		} else {
			ret[left[i].no] = abs(left[i].bloomIndex - left[i].initIndex)
		}
	}
	for i := range right {
		if right[i].bloomIndex == -1 {
			ret[right[i].no] = -1
		} else {
			ret[right[i].no] = right[i].bloomIndex - right[i].initIndex
		}
	}
	for i, v := range ret {
		fmt.Print(v)
		if i != len(ret)-1 {
			fmt.Println()
		}
	}
}

//检查这组机器人是不是全炸了,只要有一个为 -1 ，就是没有全炸
func checkIsAllBoom(robots *[]robot) bool {
	for i := 0; i < len(*robots); i++ {
		if (*robots)[i].bloomIndex == -1 {
			return false
		}
	}
	return true
}

//检查  right里面没炸的最小的nowIndex  必须大于  left里面最大的nowIndex
func check2(right *[]robot, left *[]robot) bool {
	minRight := 10000000
	for _, v := range *right {
		if v.bloomIndex == -1 {
			minRight = min(minRight, v.nowIndex)
		}
	}
	maxLeft := -1
	for _, v := range *left {
		if v.bloomIndex == -1 {
			maxLeft = max(maxLeft, v.nowIndex)
		}
	}
	return minRight < maxLeft && minRight != 10000000 && maxLeft != -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
