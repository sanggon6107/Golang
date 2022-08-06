package main

import "fmt"

// 가변 인수 함수 : ...키워드를 쓰면 가변 인수 함수를 구현할 수 있다.
// 사실 원리는 슬라이스로 변환되는 것. 즉, Arg(nums []int)와 같다.
func Arg(nums ...int) {
	for _, v := range nums {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Arg(1, 2, 3)
	Arg(1, 2, 3, 4, 5)
}
