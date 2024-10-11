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
	fmt.Print("input your score : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                  // python strip
	score, err := strconv.ParseInt(i, 16, 32) // 문자열 -> 정수 (변환할 변수, 진수크기, 비트크기)
	if score >= 60 {
		fmt.Printf("%d\n", score)
		fmt.Println("A")
	} else {
		fmt.Println("BCDF")
	}
}
