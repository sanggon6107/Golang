package main

import (
	"fmt"
	"os"
)

func TextFileIO(content string) {
	f, err := os.Create("text.txt") // 작성할 파일을 만든다.
	if err != nil {
		fmt.Println("error!")
		return
	}

	// "defer 명령어" 는, 해당 블록이 닫히기 직전에 "뒤바뀐 순서로" 실행된다. 즉, "파일을 닫습니다."가 출력되고 난 뒤 f.Close()가 실행된다.
	// 함수 끝나기 직전에 반드시 해야할 작업에 defer를 활용하면 좋다.
	defer f.Close()
	defer fmt.Println("파일을 닫습니다.")

	fmt.Fprintln(f, content) // 파일에 content 작성한다.
}

func main() {
	TextFileIO("TEST")
}
