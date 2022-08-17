package main

import "fmt"

// panic은 interface{}를 파라미터로 가진다.
// 즉, panic의 인자로는 어떠한 것도 올 수 있다. 다만 대체적으로는 Error()의 리턴값인 string이 들어간다.
// 프로그램을 종료시킨다.
// 콜스택을 순서대로 빠져나오면서 함수를 종료시키는데
// 만약 해당 함수에서 defer로 recover()를 호출시키면
// 그 밖 함수부터 다시 프로그램이 재개된다.

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
