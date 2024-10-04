package main

import (
	"fmt" //입출력 관련 라이브러리
	"reflect"
)

func main() {

	i := 13 // :=  단축연산자 ,뒤에 오는 값의 자료형을 유추해서 변수에 담아줌
	var f float32 = 12.9
	c1 := 'A' //go의 문자변수는 크기가 가변적임
	c2 := '김'
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	//fmt.Printf("%d * %f = %d\n", i, f, float(i)*f)
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
	//go 언어는 같은 자료형끼리만 연산가능

}
