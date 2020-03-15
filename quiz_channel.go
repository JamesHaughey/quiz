package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

func correctAnswer(channel chan int) {
	correct := <-channel
	correct++
	channel <- correct
}

func displayScore(correct int, questions int) {
	fmt.Printf("You got %d correct out of %d\n", correct, questions)
}

func startTimer(channel chan int, seconds int, questions int) {
	timer := time.NewTimer(time.Second * time.Duration(seconds))
	<-timer.C
	displayScore(<-channel, questions)
	os.Exit(2)
}

var csvFilePath = flag.String("csv", "problems.csv", "A csv file containing a quiz in the format of 'question,answer'.")
var timeLimit = flag.Int("time", 30, "Sets the time limit for the quiz in seconds.")

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
	// Press enter to being

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	correct := 0

	for _, problem := range problems {
		fmt.Printf("%s: ", problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			displayScore(correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
		}
	}
	displayScore(correct, len(problems))
}
