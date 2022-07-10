package main

import (
	"fmt"
	"math"
)

// Comp 는 실수형이 정확한 값을 나타내지 못해서 일어나는 오차를 보정해서 정확한 비교를 해준다.
func Comp(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

// 컴퓨터는 0과 1로 모든 수를 표현하기 때문에, 실수를 모두 정확하게 나타낼 수 없다.
// 아주 가까운 수는 표현할 수 있다. 예를 들어, 0.3은 정확하게 나타낼 수 없고 0.3보다
// 아주 미세하게 큰 값 또는 아주 미세하게 작은 값으로 표현된다. 따라서 0.3을 표현하는
// 두 가지 방식이 존재하게 되고, 상황에 따라 ==연산이 제대로 동작하지 않을 수 있다.
// 그리고 아주 미세하게 큰 값과 작은 값의 차이는 항상 마지막 비트의 차이에서 기인하는데,
// 이 차이는 math.Nextafter() 함수로 보정(?)할 수 있다.

// 이렇듯 컴퓨터가 실수 값을 표현하는 정밀도에는 한계가 있으므로, 금융 계통에서는
// math/big 패키지의 Float객체를 쓴다. Float객체를 쓰면 정밀도를 직접 조정할 수 있다.

func main() {
	fmt.Println("Hello!")
	var a float64 = 0.3
	var b float64 = 0.3
	fmt.Println(Comp(a, b))

	// 대입 연산자 인상 깊은 부분.
	var c int = 3
	var d int = 4
	fmt.Println("c : ", c, "d : ", d)
	c, d = d, c // c와 d를 서로 바꾼다.
	fmt.Println("c : ", c, "d : ", d)
}
