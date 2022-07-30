package main

import (
	"fmt"
	"sort"
)

// 슬라이스의 구조는 아래와 같다.
// type SliceHeader struct {
// 	Data uintptr // 실제 배열을 가리키는 포인터
// 	Len  int     // 길이
// 	Cap  int     // 크기
// }
// 즉, slice는 배열 그 자체라기 보다는 어떤 다른 곳에 있는 배열의 정보를 담고 있다.
// 그렇기 때문에, 함수의 파라미터(매개변수)로 슬라이스의 값타입을 받더라도 그 정보는 그대로 받는다.
// 즉, 배열의 복사본이 아니라 배열의 정보의 복사본이 만들어진다.
// 반면 Array는 정말로 배열의 복사본이 만들어진다는 것에 주의한다.

func main() {
	a := []int{1, 2, 3} // 초기화법 1 : {}로 초기화
	b := make([]int, 3) // 초기화법 2 : make(타입, 크기)함수로 초기화. 요소들의 값은 int형의 기본값인 0이 된다.
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()
	for _, v := range b {
		fmt.Print(v)
	}
	fmt.Println()

	// append()함수를 통한 요소 추가 : append는 첫 번째 인자로는 대상 슬라이스를 넣어야 한다.
	// append()는 완성품을 리턴하는데, 자기 자신에 더하기 위해서는 아래 처럼 좌변에 자기자신을 두면 된다.
	a = append(a, 4, 5, 6)
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()

	// slice의 ptr필드에 대한 이해
	c := []int{1, 2, 3, 4, 5}
	d := c // 안에 있는 포인터도 복사된다.
	d[2] = 100
	fmt.Println("c : ", c) // [1, 2, 100, 4, 5]
	fmt.Println("d : ", d) // [1, 2, 100, 4, 5]

	// slice의 len과 cap의 차이에 대한 이해
	e := []int{1, 2, 3}
	f := append(e, 4, 5)
	fmt.Println(len(e), cap(e)) // 길이(len)는 5이지만 capacity(실제 배열 길이 == 할당된 메모리)는 3
	fmt.Println(len(f), cap(f)) // 길이도 5이고 capacity도 5

	for _, v := range e {
		fmt.Print(v)
	}
	fmt.Println()
	for _, v := range f {
		fmt.Print(v)
	}
	fmt.Println()
	e = append(e, 100)

	// 상충하는 경우이므로 새로 메모리를 할당받은 듯.
	for _, v := range e {
		fmt.Print(v)
	}
	fmt.Println()
	for _, v := range f {
		fmt.Print(v)
	}
	fmt.Println()

	// 슬라이싱. 주의 : 실제 배열의 첫번째 요소에서 두번째 요소까지 가리키는 '포인터'를 가진다. 따라서,
	// g 또는 f 의 값을 바꾸게 되면 아래처럼 서로 영향을 받는다.
	g := f[1:3] // 파이썬처럼 [1:] , [:3] 과 같은 표현 가능
	fmt.Println(g)
	g[0] = 100
	fmt.Println(f)
	fmt.Println(g)

	// 슬라이싱의 복제 : 딥카피 하는 방법 1 - make
	h := []int{1, 2, 3, 4}
	aa := make([]int, len(h)) // h의 크기와 똑같은 크기의 슬라이스를 만들어서
	for i, v := range h {
		aa[i] = v // 슬라이스 안의 값을 복사한다.
	}

	// 슬라이싱의 복제 : 딥카피 하는 방법 2 - append
	bb := append([]int{}, h...) // []int{} 객체에 h의 모든 요소를 넣는다.
	// bb := append([]int{}, h[0], h[1], h[2], h[3])  과 같은 코드.

	// 슬라이싱의 복제 : 딥카피 하는 장법 3 - copy
	cc := make([]int, 2)
	ret := copy(cc, bb) // bb를 복사해서 cc에. cc의 길이는 2이다. 슬라이스 길이를 반환한다.

	// 슬라이스 정렬
	sort.Ints(cc) // sort를 import해야 한다.
}
