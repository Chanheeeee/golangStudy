package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (p PasswordError) Error() string {
	return "암호 길이가 너무 짧다."
}

// 에러를 반환하는 함수
// error : 인터페이스 타입
// 인터페이스 변수에 인터페이스 메소드를 구현한 객체 대입.
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func main() {

	err := RegisterAccount("chany", "ckslck")

	fmt.Printf("%v \n", err)
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원가입완료.")
	}

}
