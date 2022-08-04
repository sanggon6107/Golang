package main

import "fmt"

// 인터페이스는 "순수가상함수"와 비슷하다. 선언부만 제공한다. 다만 반드시 구현부를 제공해야하는 것은 아니다.
// 메소드가 struct를 받을 때, 인터페이스의 구현부를 제공한다.
// 메소드에서 인터페이스의 구현부를 따로 구현함에 따라, 다형적으로 함수를 구현할수 있게 된다.
type myInterface interface {
	myFirstMethod()
	mySecondMethod()
}

type human struct {
	name string
	age  int
}

func (h *human) myFirstMethod() {
	fmt.Println("myFirstMethod")
	fmt.Println("name : ", h.name)
	fmt.Println("age : ", h.age)
}

func main() {
	a := &human{"Kim", 20}
	a.myFirstMethod()
}
