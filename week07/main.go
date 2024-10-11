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
	score, err := strconv.ParseInt(i, 10, 32) // 문자열 -> 정수 (변환할 변수, 진수크기, 비트크기)
	var aOrNot string
	if score >= 90 {
		aOrNot = "A"
	} else {
		aOrNot = "BCDF"
	}
	fmt.Printf("%d  %s\n", score, aOrNot)
}
