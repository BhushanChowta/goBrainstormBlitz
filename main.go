package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Brainstorming!")

	f, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(questions) == 0 {
		fmt.Println("No questions found.")
		return
	}

	total := len(questions)
	used := make(map[int]bool)
	correct := 0
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < total; i++ {
		var idx int
		for {
			idx = rand.Intn(len(questions))
			if !used[idx] {
				used[idx] = true
				break
			}
			idx = rand.Intn(len(questions))
		}
		q := questions[idx]
		fmt.Printf("Q%d: %s = ", i+1, q[0])
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if strings.EqualFold(ans, q[1]) {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Printf("Wrong! Correct answer: %s\n", q[1])
		}
		fmt.Printf("Score: %d/%d\n", correct, i+1)
	}
	fmt.Printf("Final Score: %d/%d\n", correct, total)
}
