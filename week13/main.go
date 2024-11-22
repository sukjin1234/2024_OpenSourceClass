package main

import (
	"fmt"
)

func main() {
	var emptySlice []int
	//emptySlice = make([]int, 10)
	fmt.Printf("%#v\n", emptySlice)
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 4.2, 3.9}
	gpaSlice := gpa[1:4] // 4.1, 4.5, 4.2
	//gpaSlice[1] = 2.71
	gpa[2] = 5
	gpaSlice = append(gpaSlice, 9)
	fmt.Println(gpaSlice, gpa)           // slice point array
	fmt.Println(len(gpaSlice), len(gpa)) // slice point array
}
