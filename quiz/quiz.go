package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	var quizFile *os.File = fetchQuizFile(*csvFilename)
	defer quizFile.Close()

	quiz := readFile(quizFile)
	correct, incorrect := giveQuiz(quiz)

	fmt.Println("You answered ", correct, " questions correctly and ", incorrect, " incorrectly")
	fmt.Println("Your score is ", correct, "/", correct+incorrect)

}

func fetchQuizFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("Unable to open the file %s \n", path)
		os.Exit(1)
	}

	return file
}

func readFile(file *os.File) [][]string {
	reader := csv.NewReader(file)

	contents, err := reader.ReadAll()

	if err != nil {
		log.Fatal("There is a probelm parsing this file.", err)
	}

	return contents
}

func giveQuiz(quiz [][]string) (int, int) {
	var numCorrect int = 0
	var numIncorrect int = 0
	var correct bool

	for i := range quiz {

		var space string = " "
		var record string = strings.Join(quiz[i], space)
		problem, answer, found := strings.Cut(record, space)
		intAnswer, err := strconv.Atoi(answer)
		if err != nil {
			log.Fatal("There is a problem converting to integer.", err)
		}
		if !found {
			log.Fatal("This file is not formated correctly")
		}

		fmt.Println("Question", space, i+1)
		fmt.Println(problem)

		correct = getAnswer(intAnswer)

		if correct {
			numCorrect++
			fmt.Println("Correct Answer!")
		} else if !correct {
			numIncorrect++
			fmt.Println("Incorrect Answer")
			fmt.Printf("The correct answer is %d.\n", intAnswer)
		}
	}

	return numCorrect, numIncorrect
}

func getAnswer(answer int) bool {
	var correct bool
	var input int
	fmt.Scanln(&input)

	if input == answer {
		correct = true
	} else {
		correct = false
	}

	return correct
}
