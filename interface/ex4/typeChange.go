package main

import "fmt"

/*
인터페이스 변수를 다른 구체화된 타입으로 타입변환할 수 있다.
인터페이스변수.(변환할타입)

인터페이스 변수를 본래의 구체화된 타입으로 복원하려할때 주로 사용함.
*/

type Stringer interface {
	String() string
}

type Friends struct {
	Name string
}

func (f Friends) String() string {
	return fmt.Sprintf("StringName : %s", f.Name)
}

func PrintName(stringer Stringer) {
	s := stringer.(Friends)
	fmt.Printf("PrintName : %s\n", s)
	fmt.Printf("PrintName : %s\n", s.Name)
}

func main() {

	friend := Friends{"jw"}
	var stringer Stringer
	stringer = friend

	fmt.Println(stringer.String())

	PrintName(friend)

}
