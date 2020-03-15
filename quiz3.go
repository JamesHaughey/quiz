package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func askQuestion(p problem) (bool, error) {
	fmt.Printf("%s: ", p.q)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	attempt := strings.TrimSpace(input)
	if attempt == p.a {
		return true, nil
	} else {
		return false, nil
	}
}

var csvFilePath = flag.String("csv", "problems.csv", "A csv file containing a quiz in the format of 'question,answer'")

func main() {
	flag.Parse()

	csvfile, err := os.Open(*csvFilePath)
	check(err)
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	lines, err := r.ReadAll()
	check(err)
	problems := parseLines(lines)

	fmt.Println("Quiz beginning:")
	correct := 0

	for _, problem := range problems {
		result, err := askQuestion(problem)
		check(err)
		if result {
			correct++
		}
	}
	fmt.Printf("You got %d correct out of %d\n", correct, len(problems))
}
