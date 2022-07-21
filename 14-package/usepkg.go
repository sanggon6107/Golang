package main

// 같은 go 모듈 아래 있는 패키지는 go 모듈명 아래 위치하도록 해줘야 import 된다.
// 따라서, custompkg는 main패키지가 있는 go_source/14-package의 아래에 있어야 한다.
import (
	"fmt"
	"go_source/14-package/custompkg"

	"github.com/guptarohit/asciigraph"
)

func main() {
	custompkg.PrintCustom()
	fmt.Println("This is fmt!")

	data := []float64{3, 4, 5, 6, 7, 2, 3}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
