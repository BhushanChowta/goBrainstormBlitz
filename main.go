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

	fmt.Println("Press Enter to start the quiz...")
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

	total := len(questions)
	correct := 0
	reader := bufio.NewReader(os.Stdin)
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
	for i, q := range questions {
		fmt.Print("Qn ", i+1, " - ", q[0], " = ")
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if strings.EqualFold(ans, q[1]) {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Printf("Wrong! Correct answer: %s\n", q[1])
		}
	}

	fmt.Printf("Final Score: %d/%d\n", correct, total)
}
