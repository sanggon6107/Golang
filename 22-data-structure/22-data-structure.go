package main

import (
	"container/list"
	"container/ring"
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
	return que.que.Front().Value
}

func (que *Queue) Len() int {
	return que.que.Len()
}

func (que *Queue) Pop() {
	que.que.Remove(que.que.Back())
}

// 교재는 메소드가 아닌 일반 함수로 구현. Queue객체와의 결합도가 낮아져서 사용자 입장에서는 불편한 것 아닌가?
func (que *Queue) Create() {
	que.que = list.New()
}

func main() {
	// containers/list 의 list는 doubly linked list이다. 즉, 삽입과 삭제가 빈번한 경우에 적합한 불연속 메모리 구조.
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

	// queue는 FIFO
	var q Queue
	q.Create()
	q.Push(1)
	q.Push(4)
	q.Push(5)
	for q.Len() != 0 {
		fmt.Print(q.Front(), " ")
		q.Pop()
	}
	fmt.Println()

	// ring은 환형 linked list
	r := ring.New(5) // 길이를 가진다.
	ringLen := r.Len()

	for i := 0; i < ringLen; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < ringLen; j++ {
		fmt.Print(r.Value, " ")
		r = r.Next()
	}
	fmt.Println()

	// map도 있다. 해쉬맵. 파이썬의 딕셔너리.
	// Golang의 기본 타입이다. 즉, 어떠한 것도 import할 필요 없다.
	var myMap map[string]string = map[string]string{} // 교재처럼 myMap := make(map[string]string) 도 가능
	myMap["key1"] = "value1"
	myMap["key2"] = "value2"
	myMap["key3"] = "vakye3"

	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key2"])
	fmt.Println(myMap["key3"])

	for k, v := range myMap {
		fmt.Println("key : ", k, ", value : ", v) // Golang의 맵은 순서 보장하지 않는다.
	}

	delete(myMap, "key1")
	delete(myMap, "key2")
	delete(myMap, "key3")

	// 없는 키값을 찾으면 0을 반환한다. 그렇다면 이것이 진짜 0인지, 아니면 없는 값을 호출해서 0인지 어떻게 아는가?
	// -> 복수 반환에 답이 있다.
	val, ret := myMap["key4"]
	fmt.Println(val) // 0
	fmt.Println(ret) // false

	// 해시 함수가 따르는 법칙
	// 1. 같은 입력에 대해 같은 출력을 낸다.
	// 2. 다른 입력에 대해 웬만하면 다른 입력을 낸다.
	// 3. 입력값의 범위는 무한대이고 출력값의 범위는 한정되어 있다.
	// 예) 삼각함수, 나머지 연산.
}
