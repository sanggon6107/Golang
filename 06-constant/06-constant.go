package main

import "fmt"

func main() {

	var b int = 4 // 변수는 변하는 값.
	fmt.Println("b : ", b)
	b = 5
	fmt.Println("b : ", b)

	const a int = 3 // 상수는 변하지 않는 값.
	fmt.Println("a : ", a)

	// const와 괄호로 여러개의 상수를 한번에 선언할 수도 있고, iota를 활용하여 간편하게 열거값을 사용할 수도 있다.
	// iota는 각 줄의 번째수를 0부터 차례대로 넣어준다.
	const (
		red    int = iota // 0
		yellow int = iota // 1
		green  int = iota // 2
	)
	const (
		orange int = iota + 1 // 0 + 1 = 1
		blue   int = iota + 2 // 1 + 2 = 3
		indigo int = iota + 3 // 2 + 3 = 5
	)
	// 첫 번째 수와 적용되는 iota 규칙이 같다면, 아래처럼 두번째 부터는 생략할 수도 있다.
	const (
		purple int = iota + 1 // 0 + 1 = 1
		pink                  // 1 + 1 = 2
		mint                  // 2 + 2 = 3
	)

	// 타입이 없는 상수
	// 타입 없는 상수는 변수에 복사될 때 타입이 정해지기 떄문에 여러 타입에 사용되는 상숫값을 사용할 때 편리하다.
	const PI = 3.14
	var result int = PI * 100 // Go언어는 강타입 언어. PI를 float64로 선언했다면 여기서 에러가 발생했을 것이다.
	fmt.Println("PI는 int가 아니지만 100을 곱한 것이 그 값이 변수의 초기값이 되므로 에러 발생하지 않는다.")
	fmt.Println(result)

}
