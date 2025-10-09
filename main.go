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
	// This is a placeholder for the main function.
	fmt.Println("Welcome to Quiz Game!")

	f, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer f.Close()

	questions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	correctCount := 0
	for i, q := range questions {
		fmt.Println("Qn-", i+1, ":", q[0])
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if ans == q[1] {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Wrong! The correct answer is:", q[1])
		}
	}
	fmt.Println("You answered", correctCount, "out of", len(questions), "questions correctly.")
}
