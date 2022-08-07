package main

import "fmt"

func Plus(a, b int) int {
	return a + b
}

func main() {
	funcOp := Plus                      // 함수의 이름으로 초기화해버리면 함수 포인터가 된다.
	var secondFuncOp func(int, int) int // 함수의 signature가 함수 타입 변수로 쓰일 수 있다.
	secondFuncOp = Plus
	fmt.Println(funcOp(3, 4))
	fmt.Println(secondFuncOp(5, 6))
}
