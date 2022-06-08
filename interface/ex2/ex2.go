package main

import (
	"sender/fedex"
	"sender/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.KoreaPostSender{}
	SendBook("a", sender)
	SendBook("b", sender)
}
