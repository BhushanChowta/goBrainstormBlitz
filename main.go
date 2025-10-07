package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Quiz!")

	f, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	questions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	isUsed := make([]bool, len(questions))
	correctCount := 0
	streak := 0
	maxStreak := 0
	incorrectStreak := 0
	for i, q := range questions {
		if isUsed[i] {
			continue
		}
		isUsed[i] = true
		fmt.Printf("Q.%d: %s = ", i+1, q[0])
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer == q[1] {
			correctCount++
			streak++
			incorrectStreak = 0
			if streak > maxStreak {
				maxStreak = streak
			}
			fmt.Println("Correct! Current streak:", streak)
		} else {
			fmt.Println("Wrong! The correct answer is:", q[1])
			streak = 0
			incorrectStreak++
		}
	}
	fmt.Printf("Your longest correct answer streak was %d.\n", maxStreak)
	fmt.Printf("Your longest incorrect answer streak was %d.\n", incorrectStreak)
	fmt.Printf("You answered %d out of %d questions correctly.\n", correctCount, len(questions))
}
