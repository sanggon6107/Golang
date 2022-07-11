package main

import "fmt"

func main() {

	// 기본적인 for문
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 조건문만 쓸 수 있다.
	a := 3
	for a < 10 { // 또는 for ; a < 10 ;{
		fmt.Println(a)
		a++
	}

	// 무한 루프
	for true {
		fmt.Println("for true")
	}

	// 무한 루프는 true를 생략해도 된다.
	for {
		fmt.Println("for")
	}

	// break와 레이블 : break는 해당 for문만 빠져나온다. flag를 쓰거나 함수로 빼서 return해버려도 되지만,
	// 레이블을 쓰면 원하는만큼 쉽게 빠져나올 수 있다.
	// 단, 클린 코드를 위해서는 레이블은 지양할 것. (플래그나 함수화)

OuterFor: // 레이블 정의
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			if a > 5 {
				break OuterFor
			}
			fmt.Println("for")
		}
	}
}
