package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())

	n := rand.Intn(100)
	fmt.Println(n)
}
