package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// select문을 사용하면, 채널에 데이터가 들어오기를 대기하는 상황에서 데이터가 들어오지 않으면 다른 작업을 하거나,
// 여러개 채널을 동시에 대기하도록 만들 수 있다.
// select문은 아래처럼 어떤 채널에 데이터가 들어왔을 때 어떤 행동을 할 수 있는지 정할 수 있다.
func Recieve(wg *sync.WaitGroup, ch chan int, chDone chan bool) {
	for {
		select {
		case n := <-ch: // ch에 데이터가 들어오는 경우
			fmt.Println("n : ", n)
		case <-chDone: // chDone에 데이터가 들어오는 경우
			wg.Done()
			return
		}
	}
}

func main() {
	ch := make(chan int)
	chDone := make(chan bool)

	wg.Add(1)
	go Recieve(&wg, ch, chDone)
	for i := 0; i < 9; i++ {
		ch <- i // 채널이 빌때마다 넣는다.
	}
	chDone <- true
	wg.Wait()
}
