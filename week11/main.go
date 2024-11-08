package main

// package 이름과 package를 담고 있는 파일 이름이 같아야함
import "main.go/greeting"

func main() {
	greeting.AllGreeting("ll")
}
