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
	// 논리적 오류 해결
	if score < 2 { // 1 보다 큰 자연수중 1과 자신만을 약수로 가지는 수
		isPrime = false
	} else if score == 2 {
		isPrime = true
	} else if score%2 == 0 {
		isPrime = false
	} else {
		//for j := 2; j <= int(math.Sqrt(float64(score))); j++ {
		for j := 3; j*j <= score; j += 2 {
			if score%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j) //반복문 횟수 확인
		}
	}

	if isPrime { // 비교 연산 제거
		fmt.Printf("[%d 소수]", score)
	} else {
		fmt.Printf("[%d 소수 X]", score)
	}

}
