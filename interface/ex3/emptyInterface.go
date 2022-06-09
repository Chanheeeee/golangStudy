package main

import "fmt"

/*
빈 인터페이스 인수는 모든 타입을 인수로 활용할 수 있다.
*/
func PrintVal(v interface{}) {

	switch t := v.(type) {

	case int:
		fmt.Printf("v is int %d\n", int(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not Supported type: %T:%v\n", t, t)

	}

}

type student struct {
	name string
}

func main() {
	PrintVal(3)
	PrintVal("hello")
	PrintVal(student{"chany"}) //main.student:{chany}
}
