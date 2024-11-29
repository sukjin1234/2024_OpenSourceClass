package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 20241234
	student1.name = "whdo"
	student1.gpa = 4.5
	fmt.Printf("%d %s %f\n", student1.id, student1.name, student1.gpa)
	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 20241234
	student2.name = "dddd"
	student2.gpa = 4.7
	fmt.Printf("%d %s %f\n", student2.id, student2.name, student2.gpa)
}
