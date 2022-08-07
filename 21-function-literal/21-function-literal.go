package main

import "fmt"

func main() {
	// 함수 리터럴은 함수 이름 없이 구현하면 된다. 다른 언어에서 lambda라고 불리는 것과 같지만, Golang에서는 function literal이라고 불린다.

	i := 0
	f := func(a, b int) int {
		fmt.Println(i) // Capture. 참조형으로 가져옴에 주의한다. 0 출력.
		i++
		return a + b
	}

	fmt.Println(f(4, 5))
	fmt.Println(i) // 1 출력. 참조형으로 가져왔으므로 실제 i값에 반영된다.
}
