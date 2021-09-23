package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvflag := flag.String("csv", "problems.csv", "a csv file in the format of 'question', 'answer'")
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
	Correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == problem.answer {
			fmt.Println("Correct!")
			Correct++
		} else {
			fmt.Println("Wrong!")
		}
	}
	fmt.Printf("you got %d out of %d \n", Correct, len(problems))
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
