package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "quiz.csv", "a csv file contatining quiz questions & answers")
	timeLimit := flag.Int("limit", 30, "sets a time limit in seconds")
	flag.Parse()

	data, err := openCSVandReturnData(*csvFilename)

	if err != nil {
		exitMessage(fmt.Sprintf("Failed to open the CSV file: %s\nError: %v\n", *csvFilename, err))
	}

	problems := parseLinesIntoProblem(data)
	score := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

quizLoop:
	for i, problem := range problems {
		fmt.Printf("Q%d: %s = ", i+1, problem.question)

		answerChannel := make(chan string)
		go readUserInput(answerChannel)

		select {
		case <-timer.C:
			fmt.Printf("\nYour score is %d/%d\n", score, len(problems))
			break quizLoop
		case ans := <-answerChannel:
			if checkAnswer(problem, ans) {
				score++
			}
		}
	}

}

// Prints custom error message and exits the program
func exitMessage(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Opens a given CSV file and returns the data in 2D slice format
func openCSVandReturnData(filename string) ([][]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

// Parses the raw csv lines into a problem struct
func parseLinesIntoProblem(lines [][]string) []Problem {
	var result []Problem

	for _, line := range lines {
		result = append(result, Problem{question: line[0], answer: strings.TrimSpace(line[1])})
	}

	return result
}

// Checks if the user input answer matches the problem answer and returns a boolean
func checkAnswer(problem Problem, userAnswer string) bool {
	return userAnswer == problem.answer
}

// Reads the standard input from the console and passed it into a channel
func readUserInput(answerChannel chan string) {
	var input string
	fmt.Scanf("%s\n", &input)
	answerChannel <- input
}
