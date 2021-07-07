//简单的哈希
func twoSum(nums []int, target int) []int {
    m := make(map[int]int,0)
    for i,v :=range nums {
        if val,ok := m[target-v];!ok {
            m[v] = i
        }else{
            return []int{i,val}
        }
    }
    return nil
}