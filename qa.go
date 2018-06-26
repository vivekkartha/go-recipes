package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	quizzArr, err := parseCsv("questions.csv")
	if err != nil {
		panic(err)
	}
	for _, q := range quizzArr {
		fmt.Println("Question: ", q[0])
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimRight(userAnswer, "\n")
		if userAnswer == q[1] {
			fmt.Println("Correct!", q[0], "is", userAnswer)
		} else {
			fmt.Println("Wrong!", q[0], "is", q[1])
		}
	}
}

//ParseCsv reads a specified CSV file and returns a 2-D slice
func parseCsv(file string) ([][]string, error) {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
