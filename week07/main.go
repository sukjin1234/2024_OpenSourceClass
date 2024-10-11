package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("input your name:")
	name, err := input.ReadString('\n')
	//name, _ := input.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(name)
	}
}
