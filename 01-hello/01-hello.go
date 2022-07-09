// 패키지. 패키지는 코드의 묶음을 뜻한다. Go의 모든 소스 코드는 패키지에 포함되어야 한다.
// main 패키지는 프로그램의 시작점을 포함하는 특별한 패키지. main이라는 이름의 패키지는 반드시 main함수를 포함해야 한다.
// main 패키지가 아닌 패키지는 외부 패키지로 사용된다.
package main

import "fmt" // 표준 입출력 패키지를 import

// main함수
// Go언어는 함수 선언 바로 위에 함수 이름으로 시작되는 주석을 달아
// 설명을 하도록 권장하고 있다.
func main() {
	fmt.Println("Hello world!")
}

// go run hello.go

// go mod init go_source/hello
// go build
