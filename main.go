package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	csvflag := flag.String("csv", "problems.csv", "a csv file in the format of 'question', 'answer'")
	timeLimit := flag.Int("limit", 10, "time duration for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvflag)
	if err != nil {
		fmt.Printf("error opening file %s \n", err)
		os.Exit(1)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("failed to read data %s\n", err)
		os.Exit(1)
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	Correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s \n", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("you got %d out of %d \n", Correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				Correct++
			}
		}
		fmt.Printf("you got %d out of %d \n", Correct, len(problems))

	}
}

type Problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

// b:= make([][]string,3)

// b = [][]string{{"2","3"},{"2","2"}, {"3","3"} }
// parseLines(b)
