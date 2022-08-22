package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex를 쓴다고 해서 모두 좋은 것은 아니다. 동시성 프로그래밍의 이점을 얻지 못하게 될 수도 있기 때문.
// 따라서, 동시성 프로그래밍을 활용할 때는 스레드들의 역할을 나누거나, 영역을 나누어서 각 고루틴들이 상호 업무를 침범하지 않도록 적절히 뮤텍스를 활용해야한다.

var mutex sync.Mutex

func PlusMinus(num *int) {
	mutex.Lock()         // 뮤텍스를 획득하면 Unlock()할 때까지 다른 고루틴이 접근할 수 없게 된다.
	defer mutex.Unlock() // 지연으로 mutex.Unlock()으로 뮤텍스 언락.

	if *num < 0 {
		panic(fmt.Sprint("num is smaller than 0 : ", *num))
	}
	*num += 1000 // 먼저 num을 확인하고, 그다음 1000을 더한다. 하지만 뮤텍스가 없다면 두 고루틴이 동시에 num을 역참조하게 되고,
	// 결국 두 고루틴이 작동함에도 1000은 한 번만 더해지게 된다.
	time.Sleep(time.Millisecond)
	*num -= 1000
}

func main() {
	var wg sync.WaitGroup

	num := 0

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				PlusMinus(&num)
			}
		}()
		wg.Done()
	}
	wg.Wait()
}
