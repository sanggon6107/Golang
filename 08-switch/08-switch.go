package main

import "fmt"

func main() {
	a := 3
	// 각 케이스에 여러개 지정 가능.
	// c++와 거의 유사하다. 하지만 break를 쓰지 않아도 기본 해당 케이스만 실행시킨 뒤 switch문을 빠져나온다.
	switch a {
	case 1, 2, 3:
		fmt.Println("1, 2, 3")
	case 4, 5, 6:
		fmt.Println("4, 5, 6")
	default:
		fmt.Println("default")
	}

	// switch를 응용한 조건문 : switch true를 쓰면 if 처럼 조건 조사할 수 있다.
	// switch의 대상과 case가 같아지면 해당 케이스의 문장이 작동하는 원리를 활용.
	switch true {
	case a == 1 || a == 2 || a == 3:
		fmt.Println("1, 2, 3")
	case a == 4 || a == 5 || a == 6:
		fmt.Println("4, 5, 6")
	default:
		fmt.Println("default")
	}

	// if문처럼 switch init을 제공한다.
	// 만약 케이스 문장의 다음 케이스까지 실행시키고 싶다면, falltrhough를 쓸 수 있다.
	switch b := 4; b {
	case 1, 2, 3, 4:
		fmt.Println("1, 2, 3, 4")
		fallthrough
	case 5, 6, 7, 8:
		fmt.Println("5, 6, 7, 8")
	case 9:
		fmt.Println("9")
	default:
		fmt.Println("default")
	}
}
