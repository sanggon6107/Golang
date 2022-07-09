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
	// a := 1234 // 변수 선언의 다른 형태. 선언대입문이라고 부른다.

	var secondFloat float32 = 3.3
	fmt.Println(thirdInt + int(secondFloat)) // 숫자라도 각자 타입이 다르면 에러. 캐스팅 해줘야 한다. float32 -> int는 c++처럼 소수점이하 버림 처리.

	var bigInt int16 = 3456
	var smallInt int8 = int8(bigInt) // 메모리 크기 큰 것에서 작은 것으로 캐스팅하면 큰 부분의 메모리가 사라지기 때문에 값이 변경될 수 있다.

	// 컴퓨터의 숫자 표현에 대해서도 알아둔다. 부호 표현, 음의 정수의 표현(보수), 실수 표현(부동소수 : 부호+지수부+소수부)

	fmt.Println(smallInt)
}
