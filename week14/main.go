package main

import (
	"fmt"
)

type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	totalCost := 0
	for _, v := range visitors {
		totalCost += v.cost
	}
	return totalCost
}

func main() {
	var numVisitor int
	fmt.Print("방분 인원 수 입력 : ")
	fmt.Scanln(&numVisitor)
	var vs []visitor
	vs = make([]visitor, numVisitor)
	for i := 0; i < numVisitor; i++ {
		var age int
		fmt.Print("나이 입력 : ")
		fmt.Scanln(&age)
		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default:
			vs[i] = visitor{age: age, cost: 12000}
		}
	}
	fmt.Printf("total price : %d", calculateCost(vs))

}
