package main

import "fmt"

func main() {
	ages := make(map[string]int)
	var name string
	var age int
	for {
		fmt.Print("이름 입력(exit to 'q'): ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("Enter your age : ")
		fmt.Scanln(&age)
		ages[name] = age
	}
	// map에 range 사용하면 key, value로 나옴
	for name, age := range ages {
		fmt.Printf("%s is %d\n", name, age)
	}
}
