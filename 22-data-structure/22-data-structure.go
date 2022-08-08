package main

import (
	"container/list"
	"fmt"
)

// List로 queue 구현
type Queue struct {
	que *list.List
}

func (que *Queue) Push(arg interface{}) {
	que.que.PushBack(arg)
}

func (que *Queue) Front() interface{} {
	if que.que.Len() == 0 {
		return nil
	}
	return que.que.Back().Value
}

func (que *Queue) Len() int {
	return que.que.Len()
}

func (que *Queue) Pop() {
	que.que.Remove(que.que.Back())
}

func (que *Queue) Create() {
	que.que = list.New()
}

// containers/list 의 list는 doubly linked list이다. 즉, 삽입과 삭제가 빈번한 경우에 적합한 불연속 메모리 구조.
func main() {
	fmt.Println("")
	firstList := list.New()
	e1 := firstList.PushBack(4)   // 리턴값(e1)은 추가한 값을 가리킨다.
	e2 := firstList.PushFront(1)  // 리턴값(e2)은 추가한 값을 가리킨다.
	firstList.InsertBefore(3, e1) // e1 앞에 3 추가
	firstList.InsertAfter(2, e2)

	for i := firstList.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ") // 1 2 3 4
	}
	fmt.Println()

	for i := firstList.Back(); i != nil; i = i.Prev() {
		fmt.Print(i.Value, " ") // 4 3 2 1
	}
	fmt.Println()

	var q Queue
	q.Create()
	q.Push(1)
	q.Push(4)
	q.Push(5)
	for q.Len() != 0 {
		fmt.Print(q.Front(), " ")
		q.Pop()
	}
}
