package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	var num int
	attempts := 1

	fmt.Println("GUESS THE RANDOM NUMBER BETWEEN 0 AND 1000 ! !")
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &num)

	for {
		if randomNum == num {
			fmt.Println("Congratulations!! You got the secret number right")
			break
		} else {
			attempts++
			fmt.Println("Enter a number: ")
			fmt.Scanf("%d", &num)

			if attempts == 5 {
				fmt.Println("You have exhausted all attempts, try again. ")
				fmt.Printf("The secret number was %d\n", randomNum)
				break
			}
		}

	}
}
