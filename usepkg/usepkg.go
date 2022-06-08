package main

import (
	"fmt"
	"goprojectusepkg/custompkg" // 왜 여기에서 빨간줄 ..? -> 모듈만들었는데 왜 자꾸 고패쓰를 찾는건지 모르겠다

	/*
		could not import goprojectusepkg/custompkg (cannot find package "goprojectusepkg/custompkg" in any of
		/usr/local/Cellar/go/1.17.6/libexec/src/goprojectusepkg/custompkg (from $GOROOT)
		/Users/jeongchanhee/go/src/goprojectusepkg/custompkg (from $GOPATH))
	*/

	/*
		go mod tidy : go 모듈에 필요한 패키지 다운로드 -> 패키지 정보 go.mod, go.sum 파일에 적어줌
	*/
	"github.com/guptarohit/asciigraph"
)

func main() {
	fmt.Println("hi")
	custompkg.PrintCustom()

	data := []float64{3, 4, 5, 6, 9}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

}
