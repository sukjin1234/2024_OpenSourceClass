package main

import (
	"fmt"
	"reflect"
)

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 4.2, 3.9} // array
	gpaSlice := gpa[1:4]                                     // gpaSlice = slice
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))
	fmt.Println(gpa, reflect.TypeOf(gpa))
}
