package main

import (
	"bufio" // io를 담당하는 패키지
	"fmt"
	"os" // 표준 입출력 등을 가지고 있는 패키지
)

func main() {

	a := 1
	b := 2
	var c float32 = 123.456

	stdin := bufio.NewReader(os.Stdin)

	// Print() 는 출력하고 개행하지 않는다.
	fmt.Print("a : ", a, " b : ", b, " c : ", c)

	// Println() 는 출력하고 개행한다.
	fmt.Println("a : ", a, " b : ", b, " c : ", c)

	// Printf() 는 서식에 맞게 출력한다. 개행은 하지 않는다.
	fmt.Printf("a : %d, b : %d, c : %f\n", a, b, c) // 서식 문자는 %d, %s 처럼 C의 서식 문자와 비슷하다.
	// 참고로 %v는 데이터 타입에 맞춰서 기본 형태로 출력. 서식 문자를 모를 때 사용하면 유용하다.

	// 서식 설정
	fmt.Printf("%5d, %5d\n", a, b)    // 최소 너비 5로 설정. 기본은 오른쪽 정렬
	fmt.Printf("%05d, %05d\n", a, b)  // 최소 너비 5이고, 공란을 0으로 채움
	fmt.Printf("%-05d, %-5d\n", a, b) // -를 붙이면 왼쪽 정렬

	// Scan() 는 표준 입력에서 값을 입력받는다.
	n, err := fmt.Scan(&a, &b) // 입력 함수 Scan, Scanf, Scanln 은 n과 에러 코드를 반환한다.
	if err != nil {
		fmt.Println(n, err) // n으로 몇 번째 입력값이 잘못 입력되었는지 알 수 있다.
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	// Scanf() 는 서식 형태로 값을 입력받는다.
	fmt.Scanf("%d %d", &a, &b) // 서식 대로 정수 공란 정수를 받아서 각 정수를 a, b에 넣는다.
	// Scanln() 는 표준 입력에서 한 줄을 읽어서 값을 입력받는다.
	fmt.Scanln(&a, &b)
}
