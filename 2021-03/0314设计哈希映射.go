/*
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。
实现 MyHashMap 类：
MyHashMap() 用空映射初始化对象
void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更新其对应的值 value 。
int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。
*/

package main

import (
	"container/list"
)

//const base = 769

//我的做法:
type MyMyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor0314() MyMyHashMap {
	return MyMyHashMap{make([]list.List, base)}
}

/** value will always be non-negative. */
func (this *MyMyHashMap) Put(key int, value int) {
	k := this.hash(key)
	for it := this.data[k].Front(); it != nil; it = it.Next() {
		if (it.Value.([2]int))[0] == key {
			//可以直接
			//it.Value = [2]int{key, value}

			this.Remove(key) //不会直接给 it 赋值，所以选择先删除再添加
			this.data[k].PushBack([2]int{key, value})
			return
		}
	}
	this.data[k].PushBack([2]int{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyMyHashMap) Get(key int) int {
	k := this.hash(key)
	for it := this.data[k].Front(); it != nil; it = it.Next() {
		if (it.Value.([2]int))[0] == key {
			return (it.Value.([2]int))[1]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyMyHashMap) Remove(key int) {
	k := this.hash(key)
	for it := this.data[k].Front(); it != nil; it = it.Next() {
		if (it.Value.([2]int))[0] == key {
			this.data[k].Remove(it)
		}
	}
}
func (m *MyMyHashMap) hash(key int) int {
	return key % base
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

//。。。。。。。。。。。。。。。。。。。。。。。。。。。
//官方
// const base = 769

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}
