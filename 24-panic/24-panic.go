package main

import "fmt"

// panic은 interface{}를 파라미터로 가진다.
// 즉, panic의 인자로는 어떠한 것도 올 수 있다. 다만 대체적으로는 error 타입이 들어간다.
// 프로그램을 종료시킨다.
// 콜스택을 순서대로 빠져나오면서 함수를 종료시키는데
// 만약 해당 함수에서 defer로 recover()를 호출시키면
// 그 밖 함수부터 다시 프로그램이 재개된다.

// recover의 경우, 선언부를 보면 interface{}를 반환하게 되어있다. 즉, 어떠한 것도 반환 할 수 있다.
// recover()는 전파되고 있는 panic이 있는 경우에는 해당 panic 객체를 반환하고, 그렇지 않은 경우에는 nil을 반환한다.
// 따라서, Panic()의 인자로 특정 error타입을 넣고, error의 구체적인 타입(implementation)에 따라 다른recover를 하고
// 싶을 때는 다음과 같이 하면 된다.
// if r, ok := recover().(net.Error); ok {
// fmt.Println("r is net.Error Type")
// }
// 위와 같이 에러가 net.Error 타입의 에러인지 알 수 있다.

// f() starts
// return!
// panic recovered!
// main ends

func f() {
	fmt.Println("f() starts")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovered!")
		}
	}() // 함수이므로 ()로 호출
	g()
	fmt.Println("f() ends here")
}

func g() {
	fmt.Println(h(1)) // "return!"
	fmt.Println(h(0)) // panic..
	fmt.Println("g() ends here")
}

func h(num int) string {
	if num == 0 {
		panic("0!")
	}
	return "return!"
}

func main() {
	f()
	fmt.Println("main ends")
}
