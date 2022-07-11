package main

import "fmt"

func main() {
	// 쇼트 서킷(short circuit) : a == 3 연산이 true이므로, a == 4 연산은 수행하지 않는다.
	// golang의 if조건문은 기본 괄호를 쓰지 않지만, 괄호를 써서 복잡한 조건문의 순서를 정할 수는 있다.
	// ex. if 조건A || (조건B && 조건C)
	a := 3
	if a == 3 || a == 4 {
		fmt.Println("3 or 4")
	} else {
		fmt.Println("neither 3 nor 4")
	}

	var c int = 4
	var d int = 5
	// c 가 4가 아니므로, 그 뒤의 d == 8 비교 연산은 수행하지 않는다.
	if c == 3 && d == 8 {
		fmt.Println("c equals 3, and d equals 4.")
	} else {
		fmt.Println("neither c equals 3, nor d equals 4")
	}

	// c++처럼 if init을 쓸 수있다. 다중 리턴을 쓰면 여러개를 초기화 할 수도 있다.
	if y := 3; y < 5 {
		fmt.Println("y is smaller than 5")
	}
}
