package day2

import (
	"container/list"
	"testing"
)

/**
链表，TODO 什么链表？ 双链表？循环链表？
 */
func TestList(t *testing.T) {
	// 初始化
	l1 := list.New()
	// 插入
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	head := l1.PushFront(4)

	head = l1.InsertBefore(5, head)
	l1.InsertAfter(6, head)

	// 从头结点遍历
	for i := l1.Front(); i != nil; i = i.Next() {
		t.Log(i.Value)
	}
}
