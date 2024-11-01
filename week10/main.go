package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) // python strip
	score, err := strconv.Atoi(i)
	count := 0
	for j := 1; j <= score; j++ {
		if score%j == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Printf("%d 소수", score)
	} else {
		fmt.Printf("%d 소수 X", score)
	}

}
