package main

import (
	"fmt"
	"strings"
)

func main() {
	scores := 0
	userAnswer := ""
	gameOver := false

	type questionStruct struct {
		question string
		answer   string
	}
	for gameOver == false {
		q1 := questionStruct{question: "what is the capital city of kenya? ", answer: "Nairobi"}
		fmt.Println(q1.question)

		fmt.Scan(&userAnswer)
		if strings.ToLower(userAnswer) == strings.ToLower(q1.answer) {
			scores += 1
			userAnswer = ""
		} else {
			gameOver = true
		}
		fmt.Println(scores)
	}

}
