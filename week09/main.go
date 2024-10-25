package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1 //rand.Intn(0~n-1)
	fmt.Printf("%d", target)
}
