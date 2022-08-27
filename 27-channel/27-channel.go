package main

import (
	"fmt"
	"sync"
)

func Foo(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	fmt.Println("n : ", n)

	wg.Done()
}

func main() {
	// 채널이란 고루틴끼리 메세지를 전달할 수 있는 메세지 큐. 아래의 채널은 chan string이므로 스트링 타입의 메세지를 전달한다.
	var myFirstChannel chan string = make(chan string, 1) // 두 번째 인자는 버퍼의 크기. 1개까지는 저장해도 계속 프로그램이 실행된다. 2개째 넣으면 멈춤.

	myFirstChannel <- "my first message" // <- 연산자로 채널에 메세지 넣는다.

	var msg string = <-myFirstChannel // 똑같이 <- 연산자로 채널에서 메세지를 뺄 수 있다. 만약 이 때 채널에 메세지가 없다면 메세지가 들어올 때까지 기다린다.
	// 또, 이렇게 버퍼 없는 채널에 메세지를 넣어놓고 꺼내지 않으면 데드락 발생.
	fmt.Println(msg)

	// 예시 2 : 실제로 Goroutine에서 wg와 channel 활용해보기.
	var wg sync.WaitGroup
	var mySecondChannel chan int = make(chan int)

	wg.Add(1)
	go Foo(&wg, mySecondChannel)
	mySecondChannel <- 3 // 여기서 3이 들어가도 Foo함수 안의 채널에서 프린트 할 수 있다.

	wg.Wait()
}
