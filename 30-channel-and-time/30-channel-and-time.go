package main

import (
	"fmt"
	"sync"
	"time"
)

// Select문과 time 패키지를 사용하면 특정 시간마다 특정 기능을 수행하거나, 특정 시간 뒤에는 프로그램을 종료하는 등의
// 프로그램을 작성할 수 있다.
func Recieve(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // time.Tick 은, 인자로 들어온 파라미터만큼의 시간마다 Time 객체를 주는 채널을 반환한다.
	terminate := time.After(time.Second * 10) //

	for {
		select {
		case n := <-ch:
			fmt.Println("Recieved n : ", n)
			time.Sleep(time.Second)

		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("terminate")
			wg.Done()
			return
		}
	}
}
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go Recieve(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 10)
	}

	wg.Wait()
}
