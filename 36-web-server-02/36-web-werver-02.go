package main

// 쿼리 인수 (What is Query parameters?)
// http 요청을 보낼 때 필요한 인수를 쿼리 인수로 담을 수 있다.
// 쿼리 인수란 url에 붙여서 넣어주는 인수.
// ex. http://localhost?name=kim&age=20
// ?는 쿼리인수가 시작된다는 뜻. 각 인수는 key=value 형식이다.
import (
	"fmt"
	"net/http"
)

// how to parse query parameters
func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query() // query parameters can be parsed through the object "r" since http.Request objects have the information regarding http requests.
	name := values.Get("name")
	age := values.Get("age")
	if name == "" {
		name = "Anonymous"
	}
	if age == "" {
		age = "?"
	}
	fmt.Fprintf(w, "Hi %s, who is %s years old! \n", name, age)
}

type GooHandler struct {
	goo string
	bar string
}

func (g *GooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g.goo)
	fmt.Fprintln(w, g.bar)
}

func main() {

	http.HandleFunc("/", IndexPathHandler)

	http.ListenAndServe(":3000", nil)
}
