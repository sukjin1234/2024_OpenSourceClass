package main

import (
	"fmt" //입출력 관련 라이브러리
)

func main() {

	i := 55 // :=  단축연산자 ,뒤에 오는 값의 자료형을 유추해서 변수에 담아줌
	f := 12.9
	fmt.Printf("value 1 : %d, value f : %f\n", i, f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	//

}
