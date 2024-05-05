package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type QnA struct {
	Question string
	Answer   string
}

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

func Quiz(qnaSlice [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the quiz game. Please answer the following questions. Press Ctrl+C to end")
	totalQuestions := len(qnaSlice)
	correctAnswers := 0
	for _, qna := range qnaSlice {
		fmt.Printf("%s: ", qna[0])
		scanner.Scan()
		text := scanner.Text()
		if text == qna[1] {
			correctAnswers++
		}
	}
	fmt.Printf("You answered %d/%d questions correctly.\n", correctAnswers, totalQuestions)
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("csv file not specified as part of the argument")
		return
	}
	fileName := argsWithoutProg[0]
	data, err := ReadCSV(fileName)
	check(err)
	Quiz(data)
}
