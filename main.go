package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var questions = map[int]string{}
var answers = map[int]string{}

func main() {
	printLabel("Game start...")

	// initialize
	rand.Seed(time.Now().Unix())
	prepareQuestions()

	// game start
	game()
}

func prepareQuestions() {
	fp, err := os.Open("questions.txt")
	if err != nil {
		panic("Can not open the file of questions")
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Read error")
		}
		strs := strings.Split(string(line), ":")
		if len(strs) == 2 {
			answers[len(answers)] = strs[0]
			questions[len(questions)] = strs[1]
		}
	}
}

func game() {
	var r int = rand.Intn(len(questions))
	var question string = questions[r]
	var answer string = answers[r]

	fmt.Printf("%s", question)
	fmt.Println("")

	var input string = scan()
	input = strings.Trim(input, "\r\n")
	input = strings.Trim(input, "\n")

	fmt.Println("")
	printResult(input, answer)
	fmt.Println("")
	printLabel("Next question")

	game()
}

func scan() (input string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	return
}

func printLabel(label string) {
	fmt.Printf("** %s", label)
	fmt.Println("")
	fmt.Println("************************")
}

func printResult(input string, answer string) {
	if input == answer {
		fmt.Print("\x1b[32m")
		fmt.Println("Success!!")
		fmt.Print("\x1b[39m")
		return
	}

	fmt.Print("\x1b[31m")
	fmt.Println("failure…")
	fmt.Printf("Answer : %s", answer)
	fmt.Println("")
	fmt.Print("\x1b[39m")
}
