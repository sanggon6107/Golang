package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 컨텍스트는 작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할을 한다.

func TickEverySecond(wg *sync.WaitGroup, ctx context.Context) {
	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("TickEverySecond")
		case <-ctx.Done(): // context의 Done 채널에 신호가 오면 wg.Done()하고 함수 끝낸다.
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) // context 생성. 인수로 상위 컨텍스트를 넣으면 그 컨텍스트를 감싼 컨텍스트를 반환한다.
	// context.WithCancel은 컨텍스트 객체와 취소하는 함수를 반환한다.

	// ctx, _ := context.WithTimeout(context.Background(), time.Second*3) // WithTimeout 을 쓰면 3초뒤에 자동으로 ctx.Done() 채널에 시그널을 보낸다.
	// 물론 그 전에 cancel 함수를 호출하여 작업을 종료할 수도 있다.
	wg.Add(1)

	go TickEverySecond(&wg, ctx)
	time.Sleep(time.Second * 10)
	cancel() // cancel() 함수는 ctx.Done() 채널에 신호를 보낸다.
	wg.Wait()
}
