package main

import "fmt"

// 인터페이스는 "순수가상함수"와 비슷하다. 선언부만 제공한다.
// 메소드가 struct를 받을 때, 인터페이스의 구현부를 제공한다.
// 메소드에서 인터페이스의 구현부를 따로 구현함에 따라, 다형적으로 함수를 구현할수 있게 된다.
type myInterface interface {
	MyFirstMethod()
	MySecondMethod()
}

type human struct {
	name string
	age  int
}

func (h *human) MyFirstMethod() {
	fmt.Println("MyFirstMethod")
	fmt.Println("name : ", h.name)
	fmt.Println("age : ", h.age)
}

func (h *human) MySecondMethod() {
	fmt.Println("MySecondMethod")
	fmt.Println("name : ", h.name)
	fmt.Println("age : ", h.age)
}

type robot struct {
	name         string
	manufacturer string
}

func (r *robot) MyFirstMethod() {
	fmt.Println("MyFirstMethod")
	fmt.Println("name : ", r.name)
	fmt.Println("manufacturer : ", r.manufacturer)
}

func (r *robot) MySecondMethod() {
	fmt.Println("MySecondMethod")
	fmt.Println("name : ", r.name)
	fmt.Println("manufacturer : ", r.manufacturer)
}

func (r *robot) MyThirdMethod() { // 인터페이스의 메소드 이외에 다른 메소드를 포함하더라도, 인터페이스의 메소드만 충실히 구현한다면 제대로 구현한 것이다.
	fmt.Println("MyThirdMethod()")
	fmt.Println("name : ", r.name)
	fmt.Println("manufecturer : ", r.manufacturer)
}

func myFunction(inter myInterface) { // 인터페이스도 type이므로, 함수의 파라미터가 될 수 있다! 다만 pointer to interface는 안된다.
	inter.MyFirstMethod() // 구체적인 메서드 없이 인터페이스만으로도 메서드를 호출할 수 있다.
}

func main() {
	a := &human{"Kim", 20}
	a.MyFirstMethod()
	myFunction(a) // human 객체인 a는 MyFirstMethod(), MySecondMethod()를 둘다 가지고 있기 때문에 인터페이스를 구현했다고 볼 수 있다. 따라서 myInterface를 매개변수로 받는 함수에 a를 인자로 넣을 수 있다.

	b := &robot{"Robot", "Gopher"}
	b.MyFirstMethod()
	myFunction(b) // robot 객체도 MyFirstInterface를 구현했다고 볼 수 있다. 따라서 myFunction의 인자로 b가 들어갈 수 있다.
}
