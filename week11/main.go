// print prime number between two input number
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
	fmt.Print("첫 번째 정수(시작값) 입력: ")
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	a = strings.TrimSpace(a) // python strip
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("두 번째 정수(종료값) 입력: ")
	b, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	b = strings.TrimSpace(b) // python strip
	n2, err := strconv.Atoi(b)

	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) { // 비교 연산 제거
			fmt.Printf("%d ", i)
		}
	}
}
