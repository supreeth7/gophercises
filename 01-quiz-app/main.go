package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "quiz.csv", "a csv file contatining quiz questions & answers")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exitMessage(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exitMessage("Failed to parse the CSV file")
	}

	problems := parseLinesIntoProblem(lines)
	fmt.Println(problems)
}

// Prints custom error message and exits the program
func exitMessage(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Parses the raw csv lines into a problem struct
func parseLinesIntoProblem(lines [][]string) []Problem {
	result := make([]Problem, len(lines))

	for i, line := range lines {
		result[i] = Problem{question: line[0], answer: line[1]}
	}

	return result
}
