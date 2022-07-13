package main

import "fmt"

func main() {

	// array를 선언할 때, 배열의 크기에는 항상 상수가 들어가야 한다.
	// 따라서 초기치에는 [...]을 쓸 수 있지만, 타입 선언에는 [...]을 쓸 수 없다.

	// 초기화법 1
	var myFirstArray [5]int = [5]int{0, 1} // 나머지 인덱스에는 int의 기본값인 0이 초기값으로 들어간다.

	for i := 0; i < 5; i++ {
		fmt.Println(myFirstArray[i])
	}

	// 초기화법 2
	var mySecondArray = [5]int{3, 4}      // 타입 선언은 생략할 수 있다.
	for _, value := range mySecondArray { // 만약 인덱스가 필요 없다면 _로 무효화 할 수 있다.
		fmt.Println(value)
	}

	// 초기화법 3
	myThirdArray := [...]int{1, 2, 3, 4} // ...를 통해 배열 요소 개수를 생략할 수 있다.
	for i, value := range myThirdArray {
		fmt.Println("i : ", i, ", value : ", value)
	}

	// 대입 연산자를 이용한 배열의 복사
	// golang에서는 대입 연산자를 이용해서 쉽게 배열을 복사할 수 있다.
	myFirstArray = mySecondArray
	for _, value := range myFirstArray {
		fmt.Println("value : ", value)
	}

	// 다중배열 (다차원 배열)
	var myFourthArray [2][5]int = [...][5]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6}, // 마지막 컴마 있음 주의
	}
	for _, first := range myFourthArray {
		for _, second := range first {
			fmt.Print(second, " ")
		}
		fmt.Println()
	}
	// {
	// {2, 3, 4, 5, 6} } 처럼 마지막 요소와 끝나는 중괄호가 같은 줄에 있으면 마지막 컴마 생략 가능

}
