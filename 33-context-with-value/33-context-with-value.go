package main

import (
	"context"
	"fmt"
	"sync"
)

func Square(ctx context.Context) {
	if v := ctx.Value("key"); v != nil {
		n := v.(int) // ctx의 Value() 메서드는 interface{} 타입을 반환하므로 int 형식으로 바꿔서 써야 한다.
		fmt.Println("n * n =", n*n)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	ctx := context.WithValue(context.Background(), "key", 8) // key-value 형식. 값을 설정해서 다른 고루틴으로 작업을 지시할 때 외부 지시사항으로 설정할 수 있다.
	go Square(ctx)

	wg.Wait()
}
