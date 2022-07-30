package main

import "fmt"

// go언어는 class가 없고 구조체가 있다.
// 구조체는 함수를 가지지 못한다.
// 따라서 함수의 일종인 method를, 구조체의 밖에 정의한다.
// 이 때 그 메서드가 무슨 type의 메서드인지를 나타내기 위해 리시버를 쓴다.

type myType struct {
	name string
	age  int
}

func (t *myType) Info() (string, int) { // t myType이 리시버. 함수 내에서 t는 매개변수와 같은 역할을 한다. this와 비슷함.
	return t.name, t.age
}

func main() {
	object1 := &myType{"Kim", 20}
	fmt.Println(object1.name) // Kim
	fmt.Println(object1.age)  // 20
}
