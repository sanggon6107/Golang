package main

import "fmt"

type Data struct {
	name string
	age  int
}

func main() {
	var a int = 3
	var p *int = &a
	fmt.Println("address : ", p, ", value : ", *p)

	// 포인터로 포인터 변수 초기화하기 : 객체 생성 후 포인터 변수에 주소를 할당하는 것이 아니라,
	// 포인터 변수를 선언할 때 객체를 생성할 수 있다.
	// var p2 *int = &int{1}  primitive 타입에는 쓸 수 없다.
	var p3 *Data = &Data{"Kim", 20}
	fmt.Println("name : ", p3.name, ", age : ", p3.age) // -> 연산자 사용하지 않는다. 역참조 하지 않아도 . 연산자 바로 사용 가능.

	// new() 함수를 쓰면 객체 생성 가능. new()함수라고 해서 동적 할당은 아니다.
	// 단, new() 함수는 초기값을 설정할 수 없다.
	var p4 *Data = new(Data)
	p4.name = "Lee"
	p4.age = 30

	// GO언어의 경우 컴파일 타임에 탈출 검사를 해서 블록 밖에서도 쓰이는 변수는 heap에 저장한다.
	// 즉, 예를 들어 어떤 함수에서 지역 변수의 주소를 리턴한다고 하면, 이는 C++에서는 댕글링 오류가
	// 발생해야 하지만 Go에서는 힙에 메모리를 할당하므로 오류가 발생하지 않는다.
	// 또 go언어에선 스택 메모리도 계속 증가되는 동적 메모리 풀이다.

}
