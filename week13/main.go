package main

import (
	"fmt"
	"os"
	"reflect"
)

func test(strs ...string) { // ...string -> slice, like variable parameter in python
	// variable parameter should locate last parameter
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	slice := os.Args[1:] // only string
	fmt.Println(slice[1])
	fmt.Printf("%T\n", slice[1])
	slice = append(slice, "a", "b", "c", "d") // can put many argument
	fmt.Println(slice)
	test("1", "2,3,4")
	test("1")
	test("1", "2,3,4", "5,6")
	test()
}
