package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	quizzArr, err := parseCsv("questions.csv")
	if err != nil {
		panic(err)
	}
	timePtr, limitPtr := getFlags()

	go startTimer(timePtr)
	startQuiz(quizzArr, *limitPtr)
}

//startQuiz starts the quiz in a loop (until limit)
func startQuiz(quizzArr [][]string, limit int) {
	reader := bufio.NewReader(os.Stdin)
	for i, q := range quizzArr {
		if i == limit {
			os.Exit(0)
		}

		fmt.Println("\n>> What is", q[0])
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimRight(userAnswer, "\n")
		if userAnswer == q[1] {
			fmt.Println(">> Correct!", q[0], "is", userAnswer)
		} else {
			fmt.Println(">> Wrong!", q[0], "is", q[1])
		}
	}
}

//getFlags returns time and limit command-line option flag values
//time has a default of 4s and default limit is 10
func getFlags() (timePtr *int, limitPtr *int) {
	limitPtr = flag.Int("limit", 10, "--limit 10 to specify limit")
	timePtr = flag.Int("time", 4, "--time 10 to specify run-time in seconds")
	flag.Parse()
	return
}

//startTimer calls Exit() after specified time.
func startTimer(timePtr *int) {
	fmt.Println("You have", *timePtr, "s for", *timePtr, " questions")
	time.Sleep(time.Duration(*timePtr) * time.Second)
	os.Exit(0)
}

//parseCsv reads a specified CSV file and returns a 2-D slice
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
