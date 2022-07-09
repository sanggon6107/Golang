package main

import "fmt"

func main() {
	var firstVariable int = 10        // 변수 선언. 변수의 구성 요소는 이름, 값, 주소, 타입
	var secondVariable int = 20       // Go의 네이밍 컨벤션은 google c++ style guide와는 다르다.
	var firstUint uint16 = 10         // 16비트 부호 없는 정수
	var firstFloat float32 = 234.1235 // 32비트 실수
	fmt.Println(firstVariable + secondVariable)
	fmt.Println(firstUint)
	fmt.Println(firstFloat)

	var thirdInt int // 변수 초기값 생략. c++과 달리 지역 변수를 초기화하지 않으면 기본값으로 초기화 된다.
	// 기본값 : 숫자는 0 또는 0.0, 불린은 false, 문자열은 "" 등.
	a := 1234 // 변수 선언의 다른 형태.
}
