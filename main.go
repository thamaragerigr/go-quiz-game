package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		csvFileName = flag.String("csv", "problems.csv", "a file with question, answers")
		timeLimit   = flag.Int("limit", 30, "time limit for the quiz in seconds")
	)

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

	// we're declaring the timer here so the timer
	// is initiated after all the set up work for
	// the quiz is done
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnswers := 0

	// this is a label to break the loop
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerChanel := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			// whenever we get an answer we send it to the answerChanel
			answerChanel <- answer
		}()

  // selects are specific to goroutines
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerChanel:
			if answer == p.answers {
				fmt.Println("Correct!")
				correctAnswers++
			}
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
			answers: strings.TrimSpace(line[1]),
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
