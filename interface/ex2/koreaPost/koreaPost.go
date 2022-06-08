package koreaPost

import "fmt"

type KoreaPostSender struct {
}

// interface의 메소드 구현
func (f KoreaPostSender) Send(parcel string) {
	fmt.Printf("KoreaPost Sends %v parcel\n", parcel)
}
