package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int  // zero value
	var b bool //선언 후 자동 초기화
	var s string

	fmt.Println(f, b, s, i) //문자열값이 없으므로 출력 X
	fmt.Printf("%f %t %s %d\n", f, b, s, i)
	f = 2.7
	i = 3
	fmt.Print(f > float64(i), "\n") //comparsion (true/false)
}
