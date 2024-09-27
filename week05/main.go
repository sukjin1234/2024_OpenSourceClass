package main

import "fmt" //입출력 관련 라이브러리

func main() {
	// var i int = 55
	// var i int
	// i = 55

	i := 55 // :=  뒤에 오는 값의 자료형을 유추해서 변수에 담아줌

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i :", i)
}
