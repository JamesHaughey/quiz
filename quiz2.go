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

func askQuestion(question string, answer string) (bool, error) {
	fmt.Printf("%s: ", question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	attempt := strings.TrimSpace(input)
	if attempt == answer {
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

	fmt.Println("Quiz beginning:")
	correct := 0

	for _, line := range lines {
		question := line[0]
		answer := line[1]
		result, err := askQuestion(question, answer)
		check(err)
		if result {
			correct += 1
		}
	}
	fmt.Printf("You got %d correct out of %d\n", correct, len(lines))
}
