package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1 //rand.Intn(0~n-1)
	fmt.Printf("%d\n", answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번 남음, input your score : ", 3-guesses)
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		score, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(score)

		if answer == score {
			fmt.Println("정답")
		} else if answer > score {
			fmt.Println("Low")
		} else if answer < score {
			fmt.Println("High")
		}
	}

}
