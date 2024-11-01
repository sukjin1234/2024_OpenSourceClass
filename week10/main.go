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

	//count := 0
	var isPrime bool = true //가독성, 저장공간크기 개선
	for j := 2; j < score; j++ {
		if score%j == 0 {
			isPrime = false
		}
	}

	if isPrime { // 비교 연산 제거
		fmt.Printf("%d 소수", score)
	} else {
		fmt.Printf("%d 소수 X", score)
	}

}
