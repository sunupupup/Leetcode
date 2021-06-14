package main

import "container/list"
const base = 769

type MyHashSet struct {
    data []list.List
}

func (this *MyHashSet)hash(key int)int{
    return key%base
}

/** Initialize your data structure here. */
func Constructor0313() MyHashSet {
    return MyHashSet{make([]list.List,base)}
}


func (this *MyHashSet) Add(key int)  {
    if !this.Contains(key) {
        k := this.hash(key)
        this.data[k].PushBack(key)
    }
}


func (this *MyHashSet) Remove(key int)  {
    k := this.hash(key)
    for item:=this.data[k].Front();item!=nil;item=item.Next(){
        if item.Value.(int)==key {
            this.data[k].Remove(item)
        }
    }
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    k := this.hash(key)
    for item:=this.data[k].Front();item!=nil;item=item.Next(){
        if item.Value.(int)==key {
            return true
        }
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */