package main

import (
	"bufio"
	"fmt"
)

var stdin = bufio.Wri

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err

}

func main() {
	for {
		fmt.Printf("숫자를 입력하세요.")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			fmt.Println("입력하신 숫자는 ", n, "입니다.")
		}
	}
}
