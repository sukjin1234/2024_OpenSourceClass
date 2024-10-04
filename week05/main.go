package main

import (
	"fmt" //입출력 관련 라이브러리
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int = 55
	// var i int
	// i = 55

	i := 55 // :=  단축연산자 ,뒤에 오는 값의 자료형을 유추해서 변수에 담아줌

	var f float32 = 4.30
	//f := 4.30 //float64 에 할당됨

	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))

	fmt.Printf("%s\n", strings.Title("inhatc class"))
	fmt.Println(math.Ceil(3.99))

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i :", i)
}
