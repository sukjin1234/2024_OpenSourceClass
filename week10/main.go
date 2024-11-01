package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool { //함수화
	if n < 2 { // 1 보다 큰 자연수중 1과 자신만을 약수로 가지는 수
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		//for j := 2; j <= int(math.Sqrt(float64(n))); j++ {
		for j := 3; j*j <= n; j += 2 {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j) //반복문 횟수 확인
		}
	}
	return true
}

func main() {
	fmt.Print("정수 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) // python strip
	score, err := strconv.Atoi(i)

	if isPrime(score) { // 비교 연산 제거
		fmt.Printf("[%d 소수]", score)
	} else {
		fmt.Printf("[%d 소수 X]", score)
	}

}
