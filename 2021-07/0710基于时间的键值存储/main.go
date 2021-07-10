/*
981. 基于时间的键值存储
创建一个基于时间的键值存储类 TimeMap，它支持下面两个操作：

1. set(string key, string value, int timestamp)
存储键 key、值 value，以及给定的时间戳 timestamp。

2. get(string key, int timestamp)
返回先前调用 set(key, value, timestamp_prev) 所存储的值，其中 timestamp_prev <= timestamp。
如果有多个这样的值，则返回对应最大的  timestamp_prev 的那个值。
如果没有值，则返回空字符串（""）。

*/

package main

type T struct {
	value     string
	timestamp int
}
type TimeMap struct {
	m map[string][]T
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]T, 0),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.m[key] == nil {
		this.m[key] = make([]T, 0)
	}
	this.m[key] = append(this.m[key], T{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	//找到key对应的value，且这个value的时间戳 <= timestamp,返回存在的最大的一个时间戳

	if v, ok := this.m[key]; !ok {
		return ""
	} else {
		index := helper(v, timestamp)
		if index == -1 {
			return ""
		} else {
			return v[index].value
		}
	}
	return ""
}

//在切片中找到小于等于target的最大的一个数,返回下标  可以直接搜索，也可以二分
func helper(s []T, t int) int {
	if s == nil {
		return -1
	}
	if s[0].timestamp > t {
		return -1
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i].timestamp <= t {
			return i
		}
	}
	return -1
}

//二分

//在切片中找到小于等于target的最大的一个数,返回下标
func helper2(s []T, t int) int {
	if s == nil {
		return -1
	}
	if s[0].timestamp > t {
		return -1
	}
	//找到小于t的最大的一个数
	l, r := 0, len(s)-1
	for r > l {
		mid := (r + l + 1) >> 1
		if s[mid].timestamp > t {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

//也可以理解位找到第一个大于t的数，下标为index，返回index-1
//在切片中找到小于等于target的最大的一个数,返回下标
func helper3(s []T, t int) int {
	if s == nil {
		return -1
	}
	if s[0].timestamp > t {
		return -1
	}

	//理解为找到第一个大于t的数，然后返回index-1
	l, r := 0, len(s)-1
	for r >= l {
		mid := (r + l) >> 1
		if s[mid].timestamp <= t {
			l = mid + 1
		} else {
			if r == l {
				return r - 1 //防止死循环
			}
			r = mid
		}
	}
	return l - 1
}
