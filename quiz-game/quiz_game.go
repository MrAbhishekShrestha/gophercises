package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadCSV(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func Quiz(qnaSlice [][]string, maxSeconds int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the quiz game. Please answer the following questions. Press Ctrl+C to end")
	correctAnswers := 0
	timer := time.NewTimer(time.Duration(maxSeconds) * time.Second)
	answerCh := make(chan string)

	for _, qna := range qnaSlice {
		fmt.Printf("%s: ", qna[0])
		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}()
		select {
		case <-timer.C:
			fmt.Println()
			fmt.Printf("You answered %d/%d questions correctly.\n", correctAnswers, len(qnaSlice))
			return
		case ans := <-answerCh:
			if ans == qna[1] {
				correctAnswers++
			}
		}
	}
	fmt.Printf("You answered %d/%d questions correctly.\n", correctAnswers, len(qnaSlice))
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("csv file not specified as part of the argument")
		return
	}
	fileName := argsWithoutProg[0]
	const DEFAULT_TIMELIMIT_SECONDS = 15
	maxSeconds := func() int {
		sec := DEFAULT_TIMELIMIT_SECONDS
		if len(argsWithoutProg) == 2 {
			timeArg, err2 := strconv.Atoi(argsWithoutProg[1])
			check(err2)
			sec = timeArg
		}
		return sec
	}()
	data, err := ReadCSV(fileName)
	check(err)
	Quiz(data, maxSeconds)
}
