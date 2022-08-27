package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Receive(wg *sync.WaitGroup, ch chan int) {

	for i := range ch { // 채널에 값이 들어올 때마다 Println. 만약 Close(ch) 하지 않으면 해당 for문은 ch에 값이 들어올 때까지 무한 반복한다.
		fmt.Println("i : ", i)
	}
	wg.Done()
}

func main() {
	ch := make(chan int)

	wg.Add(1)
	go Receive(&wg, ch)
	for i := 0; i < 9; i++ {
		ch <- i // 채널이 빌때마다 넣는다.
	}
	close(ch) // 위 for문 끝나고 채널을 닫기 때문에, Recieve 함수 내의 채널이 더이상 데이터가 들어올 때까지 기다리지 않고 for문을 빠져나가 Done()을 호출한다.
	wg.Wait()
}
