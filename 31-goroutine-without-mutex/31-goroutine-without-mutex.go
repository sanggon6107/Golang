package main

import (
	"fmt"
	"sync"
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

func InstallWindow(chIn, chOut chan *Car, purchaseOrder *PurchaseOrder) {
	fmt.Println("InstallWindow")
	for data := range chIn {
		data.window = purchaseOrder.window
		chOut <- data
	}
}

func InstallTire(carCh chan *Car, purchaseOrder *PurchaseOrder) {
	fmt.Println("InstallTire")
}

func PaintCar(carCh chan *Car, purchaseOrder *PurchaseOrder) {
	fmt.Println("PaintCar")
}

func main() {
	var wg sync.WaitGroup

	wg.Wait()

}
