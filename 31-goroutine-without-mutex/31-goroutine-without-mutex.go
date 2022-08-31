package main

import (
	"fmt"
	"sync"
	"time"
)

// 채널을 만들면 뮤텍스 없이도 Go routine의 역할을 안전하게 나눌 수 있다.
// 문 부착 -> 타이어 부착 -> 차 도색 순으로 자동차가 완성된다고 하면, 고 루틴으로 각각의 기능을 나누고 채널로 다음 공정으로
// 넘기는 식으로 구현할 수 있다. 즉, 컨베이어벨트 식의 작업이 가능해진다. 각 공정이 1초씩 걸린다면 처음의 차만 3초 걸리고
// 그 뒤부터는 1초에 한대가 완성된다.

type Car struct {
	color  string
	tire   string
	window string
}

type PurchaseOrder struct {
	color  string
	tire   string
	window string
}

func InstallWindow(chOut chan *Car, wg *sync.WaitGroup, purchaseOrder *PurchaseOrder) {
	tick := time.Tick(time.Second)
	after := time.After(time.Second * 10)

	for {
		select {
		case <-tick:
			fmt.Println("InstallWindow")
			car := &Car{}
			car.window = purchaseOrder.window
			chOut <- car
		case <-after:
			close(chOut)
			wg.Done()
			return
		}
	}
}

func InstallTire(chIn, chOut chan *Car, wg *sync.WaitGroup, purchaseOrder *PurchaseOrder) {
	for car := range chIn {
		fmt.Println("InstallTire")
		car.tire = purchaseOrder.tire
		time.Sleep(time.Second)
		chOut <- car
	}
	wg.Done()
	close(chOut)
}

func PaintCar(chIn chan *Car, wg *sync.WaitGroup, purchaseOrder *PurchaseOrder) {
	for car := range chIn {
		fmt.Println("PaintCar")
		car.color = purchaseOrder.color
		time.Sleep(time.Second)

	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	po := &PurchaseOrder{"Blue", "big tire", "small window"}
	chTire := make(chan *Car)
	chPaint := make(chan *Car)
	go InstallWindow(chTire, &wg, po)
	go InstallTire(chTire, chPaint, &wg, po)
	go PaintCar(chPaint, &wg, po)

	wg.Wait()
	fmt.Println("close factory")
}
