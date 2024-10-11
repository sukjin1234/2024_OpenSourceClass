package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	//i, err := r.ReadString('\n')
	fmt.Print("input your name:")
	name, _ := input.ReadString('\n')
	fmt.Print(name)
}
