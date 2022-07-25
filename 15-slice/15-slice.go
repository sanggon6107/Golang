package main

import "fmt"

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
}
