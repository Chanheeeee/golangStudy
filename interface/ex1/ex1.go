package main

import "fmt"

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
