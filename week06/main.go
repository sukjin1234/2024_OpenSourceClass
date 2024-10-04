package main

import (
	"fmt" //입출력 관련 라이브러리
	"reflect"
)

func main() {

	i := 13 // :=  단축연산자 ,뒤에 오는 값의 자료형을 유추해서 변수에 담아줌
	f := 12.9
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	//fmt.Printf("%d * %f = %d\n", i, f, float(i)*f)
	fmt.Println(reflect.TypeOf(i))
	//go 언어는 같은 자료형끼리만 연산가능

}
