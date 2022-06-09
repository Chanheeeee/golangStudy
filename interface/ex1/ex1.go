package main

import "fmt"

/*
Friends라는 객체는 인터페이스의 String() 메서드를 포함하고 있기에
Stringer 인터페이스로 사용될 수 있다.

Sprintf : 화면에 출력하는 것이 아닌, string타입으로 반환
*/

type Stringer interface {
	String() string
}

type Friends struct {
	Name string
}

func (f Friends) String() string {
	return fmt.Sprintf("name : %s", f.Name)
}

func main() {

	friend := Friends{"jw"}
	var stringer Stringer
	stringer = friend

	fmt.Println(stringer.String())

}
