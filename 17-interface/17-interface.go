package main

import "fmt"

// 인터페이스는 "순수가상함수"와 비슷하다. 선언부만 제공한다.
// 메소드가 struct를 받을 때, 인터페이스의 구현부를 제공한다.
// 메소드에서 인터페이스의 구현부를 따로 구현함에 따라, 다형적으로 함수를 구현할수 있게 된다.

// 덕타이핑
// Golang에서는 struct가 메소드를 통해 인터페이스를 구현할 때, "어떤 인터페이스를 구현하였는가?"를 명시할 필요가 없다.(아래 코드 참조)
// 따라서 사용자는 사용자 마음대로 인터페이스를 구현할 수 있고, 각종 패키지 개발자들 역시 마음대로 함수를 구현할 수 있다. 그리고 사용자는 인터페이스에
// 필요한 메소드들의 집합만 넣으면, 서로 다른 개발자들에게서 개발된 패키지들이라고 하더라도 메소드 이름만 같을 때는 인터페이스를 받는 함수를
// 쉽게 구현할 수 있다. 즉, 이처럼 인터페이스는 '패키지 개발자'들과 '패키지 이용자'들을 매개하는 역할을 한다.

type myInterface interface {
	MyFirstMethod()
	MySecondMethod()
}

type mySecondInterface interface {
	MyFirstMethod()
	Second()
}

// interface 가 다른 interface를 포함할 수 있다.
type myFirstSecondInterface interface {
	myInterface       // myInterface의 메서드 집합을 포함한다.
	mySecondInterface // mySecondInterface의 메서드 집합을 포함한다.
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

func (r *robot) MyThirdMethod() { // 인터페이스의 메소드 이외에 다른 메소드를 포함하더라도, 인터페이스의 메소드만 충실히 구현한다면 제대로 구현한 것이다. 이것을 '덕타이핑'이라고 부른다.
	fmt.Println("MyThirdMethod()")
	fmt.Println("name : ", r.name)
	fmt.Println("manufecturer : ", r.manufacturer)
}

// MyFunction을 호출할 때 인자로는, myInterface를 구현한 struct의 인스턴스를 넣으면 된다. 즉, myInterface를 충실히 구현한 어떠한 객체도 여기에 넣을 수 있다.
func MyFunction(inter myInterface) { // 인터페이스도 type이므로, 함수의 파라미터가 될 수 있다! 다만 pointer to interface는 안된다.
	inter.MyFirstMethod() // 구체적인 메서드 없이 인터페이스만으로도 메서드를 호출할 수 있다.
}

// 빈 인터페이스의 활용 : 인터페이스가 비어있으면, 모든 객체가 빈 인터페이스의 구현이 될 수 있다. 어떤 객체라도 받고 싶다면 텅 빈 인터페이스를 활용할 수 있다.
func MySecondFunction(inter interface{}) {
	switch t := inter.(type) { // 어떤 객체가 나오든,  그 객체의 타입을 조사해서 조건 분기 구현
	case int:
		fmt.Println("int : ", t)
	case float32:
		fmt.Println("float32 : ", t)
	case float64:
		fmt.Println("float64 :", t)
	default:
		fmt.Println("??????")
	}
}

// 참고 : 인터페이스의 변환
// concrete type이 한번에 여러개의 인터페이스를 구현하였을 때,
// t, ok := 객체.(변환할 인터페이스)로 인터페이스를 바꿀 수 있다. ok에는 변환의 성공 여부를 리턴받을 수 있다.

func main() {
	a := &human{"Kim", 20}
	a.MyFirstMethod()
	MyFunction(a) // human 객체인 a는 MyFirstMethod(), MySecondMethod()를 둘다 가지고 있기 때문에 인터페이스를 구현했다고 볼 수 있다. 따라서 myInterface를 매개변수로 받는 함수에 a를 인자로 넣을 수 있다.

	b := &robot{"Robot", "Gopher"}
	b.MyFirstMethod()
	MyFunction(b) // robot 객체도 MyFirstInterface를 구현했다고 볼 수 있다. 따라서 MyFunction의 인자로 b가 들어갈 수 있다.
}
