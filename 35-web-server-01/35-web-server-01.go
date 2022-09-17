package main

import (
	"fmt"
	"net/http"
)

// 핸들러란 http 요청 url이 수신되었을 때 그것을 처리하는 함수 또는 객체이다.
func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Index!")
}

// 객체 핸들러
type GooHandler struct {
	goo string
	bar string
}

// http.handler 인터페이스를 구현해야 한다. The object needs to implement the interface called "http.handler", which has a method called "ServeHTTP".
func (g *GooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g.goo)
	fmt.Fprintln(w, g.bar)
}

func main() {
	// 인덱스 페이지 핸들러 함수를 등록한다. "/" 경로에 해당하는 http 요청을 수신할 때 IndexPathHandler 함수를 호출한다.
	http.HandleFunc("/", IndexPathHandler)

	http.HandleFunc("/Foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Foo!")
	})

	http.Handle("/Goo", &GooHandler{"AAA", "BBB"})

	// 서버 시작. 3000번 포트에서 기다린다.
	http.ListenAndServe(":3000", nil)
}
