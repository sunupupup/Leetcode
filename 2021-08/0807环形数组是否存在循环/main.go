/*
存在一个不含 0 的 环形 数组 nums ，每个 nums[i] 都表示位于下标 i 的角色应该向前或向后移动的下标个数：

如果 nums[i] 是正数，向前 移动 nums[i] 步
如果 nums[i] 是负数，向后 移动 nums[i] 步
因为数组是 环形 的，所以可以假设从最后一个元素向前移动一步会到达第一个元素，而第一个元素向后移动一步会到达最后一个元素。

数组中的 循环 由长度为 k 的下标序列 seq ：
1. 遵循上述移动规则将导致重复下标序列 seq[0] -> seq[1] -> ... -> seq[k - 1] -> seq[0] -> ...
2. 所有 nums[seq[j]] 应当不是 全正 就是 全负
3. k > 1
如果 nums 中存在循环，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [2,-1,1,2,2]
输出：true
解释：存在循环，按下标 0 -> 2 -> 3 -> 0 。循环长度为 3 。

示例 2：

输入：nums = [-1,2]
输出：false
解释：按下标 1 -> 1 -> 1 ... 的运动无法构成循环，因为循环的长度为 1 。根据定义，循环的长度必须大于 1 。

*/

//官方：快慢指针
//从每个节点出发，快指针每次移动两次，慢指针每次移动一次
func circularArrayLoop(nums []int) bool {
	n := len(nums)

	mask := make([]int, n)
	next := map[int]int{}
	for i, v := range nums {
		if abs(v)%n == 0 {
			mask[i] = -1 //表示这个点，没法构成循环数组
		}
		next[i] = (1000*n + i + v) % n
	}
	fmt.Println(next)
	for i := range nums {
		//快慢指针
		if mask[i] == -1 {
			continue
		}
		slow := i
		fast := i
		for {

			//判断slow指针的移动是否合法
			if nums[slow]*nums[next[slow]] < 0 {
				break
			}
			//判断fast指针的移动是否合法
			if nums[fast]*nums[next[fast]] < 0 || nums[fast]*nums[next[next[fast]]] < 0 {
				break
			}

			slow = next[slow]
			fast = next[next[fast]]
			if mask[slow] == -1 || mask[fast] == -1 {
				break
			}
			if slow == fast {
				return true
			}

		}
		mask[i] = -1
	}

	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

//我的做法：从每个节点深搜过去，查看是否有循环，没有循环就将节点标记为-1，后续再访问到就直接false
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	//if nums[i] == n { continue   }
	//先构建map   key:当前下标   value:下一个下标
	mask := make([]int, n)
	m := map[int]int{}
	for i, v := range nums {
		if abs(v)%n == 0 {
			mask[i] = -1 //表示这个点，没法构成循环数组
		}
		m[i] = (1000*n + i + v) % n //这里需要对往后移的下标进行判断
	}

	fmt.Println(m)

	var helper func(int, *map[int]struct{}) bool
	helper = func(now int, temp *map[int]struct{}) bool {
		if mask[now] == -1 {
			return false
		}
		if _, ok := (*temp)[now]; ok {
			return true
		} else {
			(*temp)[now] = struct{}{}
		}

		next := m[now]

		if mask[next] == -1 || (nums[now]*nums[next]) < 0 {
			return false
		}

		return helper(next, temp)
	}

	for i := range nums {
		//从这个点一路走过去，全标记为灰
		//mm统计这一路的节点
		mm := map[int]struct{}{}

		if helper(i, &mm) {
			fmt.Println(i)

			return true
		} else {
			mask[i] = -1
		}

	}

	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

