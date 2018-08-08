package lru

import (
	"container/list"
	"sync"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type cacheItem struct {
	Key string
	Val interface{}
}

type LRU struct {
	maxNum	int
	curNum	int
	mutex sync.Mutex
	data  *list.List

}

//添加数据
func (lru *LRU) Add(key string, value interface{}) error {
	//判断key是否存在
	if e, _ := lru.Exist(key); e {
		return errors.New(key + "已存在");
	}
	//判断当前存储数量与最大存储数量
	if lru.maxNum == lru.curNum {
		//链表已满，则删除链表尾部元素
		lru.Clear()
	}
	lru.mutex.Lock()
	lru.curNum++
	//json序列化数据
	data, _ := json.Marshal(cacheItem{key, value})
	//把数据保存到链表头部
	lru.data.PushFront(data)
	lru.mutex.Unlock()
	return nil
}

//设置数据
func (l *LRU) Set(key string, value interface{}) error {
	e, item := l.Exist(key)
	if !e {
		return l.Add(key, value)
	}
	l.mutex.Lock()
	data, _ := json.Marshal(cacheItem{key, value});
	//设置链表元素数据
	item.Value = data
	l.mutex.Unlock()
	return nil
}

// 清理数据
func (lru *LRU) Clear() interface{} {
	lru.mutex.Lock()
	lru.curNum--
	//删除链表尾部元素
	v := lru.data.Remove(lru.data.Back())
	lru.mutex.Unlock()
	return v
}

// 获取数据
func (lru *LRU) Get(key string) interface{} {
	e, item := lru.Exist(key)
	if !e {
		return errors.New(key + " not exist")
	}
	lru.mutex.Lock()
	lru.data.MoveToFront(item)
	lru.mutex.Unlock()
	var data cacheItem;
	json.Unmarshal(item.Value.([]byte), &data)
	return data.Val
}

// 删除数据
func (lru *LRU) Del(key string) error{
	e, item := lru.Exist(key)
	if !e {
		return errors.New(key + " not exist")
	}
	lru.mutex.Lock()
	lru.curNum--
	lru.data.Remove(item)
	lru.mutex.Unlock()
	return nil

}

// 判断key是否存在
func (lru *LRU) Exist(key string) (bool, *list.Element){
	var data cacheItem
	for v := lru.data.Front(); v != nil; v= v.Next() {
		json.Unmarshal(v.Value.([]byte), &data)
		if key == data.Key {
			return true, v
		}
	}
	return false, nil
}

func (lru *LRU) IsExist(key string) bool{
	var data cacheItem
	for v := lru.data.Front(); v != nil; v= v.Next() {
		json.Unmarshal(v.Value.([]byte), &data)
		if key == data.Key {
			return true
		}
	}
	return false
}

// 返回长度
func (lru *LRU) Len() int {
	return lru.curNum
}

// 打印链表
func (lru *LRU) Print() {
	var data cacheItem
	for v := lru.data.Front(); v != nil; v = v.Next() {
		json.Unmarshal(v.Value.([]byte), &data);
		fmt.Println("key:", data.Key, " value:", data.Val);
	}
}


// 创建LRUCache
func NewLRUCache(maxNum int) *LRU {
	return &LRU{
		maxNum: maxNum,
		curNum: 0,
		mutex: sync.Mutex{},
		data: list.New(),
	}
}
