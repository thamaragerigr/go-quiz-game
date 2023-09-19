package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a file with question, answers")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	problems := parseLines(lines)

	correctAnswers := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answers {
			fmt.Println("Correct!")
			correctAnswers++
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correctAnswers, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
   // make sure all answers are actually answerable
			answers:  strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answers  string
}

func exit(msg string) {
	fmt.Println((msg))
	os.Exit(1)
}
