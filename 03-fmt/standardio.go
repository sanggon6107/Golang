package main

import "fmt"

func main() {

	a := 1
	b := 2
	var c float32 = 123.456

	// Print() 는 출력하고 개행하지 않는다.
	fmt.Print("a : ", a, " b : ", b, " c : ", c)

	// Println() 는 출력하고 개행한다.
	fmt.Println("a : ", a, " b : ", b, " c : ", c)

	// Printf() 는 서식에 맞게 출력한다. 개행은 하지 않는다.
	fmt.Printf("a : %d, b : %d, c : %f\n", a, b, c) // 서식 문자는 %d, %s 처럼 C의 서식 문자와 비슷하다.
	// 참고로 %v는 데이터 타입에 맞춰서 기본 형태로 출력. 서식 문자를 모를 때 사용하면 유용하다.
}
