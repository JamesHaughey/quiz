package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
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

func main() {
	csvfile, err := os.Open("problems.csv")
	check(err)
	defer csvfile.Close()
	r := csv.NewReader(csvfile)

	fmt.Println("Quiz beginning:")
	questions := 0
	correct := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		questions += 1
		question := record[0]
		answer := record[1]
		result, err := askQuestion(question, answer)
		check(err)
		if result {
			correct += 1
		}
		fmt.Println(questions, correct)
	}

	fmt.Printf("You got %d correct out of %d\n", correct, questions)
}
