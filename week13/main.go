package main

import (
	"fmt"
)

func main() {
	var emptySlice []int // slice can get nil data
	//emptySlice = make([]int, 5) // []int{0,0,0,0,0}
	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, 1)
	}
	fmt.Printf("%#v\n", emptySlice)
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 4.2, 3.9}
	gpaSlice := gpa[1:4] // 4.1, 4.5, 4.2
	//gpaSlice[1] = 2.71
	gpa[2] = 5
	gpaSlice = append(gpaSlice, 9)
	fmt.Println(gpaSlice, gpa)           // slice point array
	fmt.Println(len(gpaSlice), len(gpa)) // slice point array
}
