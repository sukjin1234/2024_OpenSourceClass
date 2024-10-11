package main

import (
	"fmt"
	"strings"
)

func main() {
	var broken string = "1 . 2 . 3 . 4"
	r := strings.NewReplacer(".", ",")
	fmt.Println(r.Replace(broken))
}
