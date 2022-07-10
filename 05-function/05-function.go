package main

import "fmt"

// 함수가 외부에서 받아오는 값을 파라미터 또는 매개 변수라고 부른다.
// Plus 함수는 파라미터 두 개를 더한 값을 반환한다.
func Plus(a int, b int) int {
	return a + b
}

// 멀티 반환 함수 : 반환값이 여러개이다.
// PlusMinus 함수는 파라미터를 더한 값과 뺀 값을 리턴한다.
func PlusMinus(a int, b int) (int, int) {
	return a + b, a - b
}

// 리턴값의 이름을 함수의 선언부에서 지정하기
// Divide 함수는 파라미터를 나눈 값을 리턴한다.
func Devide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 반환값을 지정하지 않은 return문.
	} else {
		result = a / b
		success = true
		return
	}
}

func main() {
	firstArg := 4
	secondArg := 5
	fmt.Println(Plus(firstArg, secondArg)) // 함수를 호출할 때 입력하는 값을 argument 또는 인수라고 부른다.
	// 함수가 호출될 때 인수(argument)가 파라미터로 복사된다.
}
