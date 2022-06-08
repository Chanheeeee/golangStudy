package main

import (
	"sender/fedex"
	"sender/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.KoreaPostSender{}
	SendBook("a", koreaPostSender)
	fedexSender := &fedex.FedexSender{}
	SendBook("b", fedexSender)
}
