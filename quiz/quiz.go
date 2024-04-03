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

	//var quizPath string = quizLocation()
	//var quizFile *os.File = fetchQuizFile(quizPath)
	var quizFile *os.File = fetchQuizFile(*csvFilename)
	defer quizFile.Close()

	quiz := readFile(quizFile)
	correct, incorrect := giveQuiz(quiz)

	fmt.Println("You answered ", correct, " questions correctly and ", incorrect, " incorrectly")
	fmt.Println("Your score is ", correct, "/", correct+incorrect)

}

func quizLocation() string {
	var userPath string

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get working directory.", err)
	}
	fmt.Println("The current working directory is: ", workingDir)
	fmt.Println("Please enter file path to quiz")
	fmt.Scanln(&userPath)

	return userPath
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
	var correct int = 0
	var incorrect int = 0

	for i := range quiz {
		var userIn int
		var space string = " "
		var record string = strings.Join(quiz[i], space)
		problem, answer, found := strings.Cut(record, space)
		intAnswer, err := strconv.Atoi(answer)
		if err != nil {
			log.Fatal("There is a problem converting to integer.", err)
		}
		if found != true {
			log.Fatal("This file is not formated correctly")
		}

		fmt.Println("Question", space, i+1)
		fmt.Println(problem)
		fmt.Scanln(&userIn)

		if userIn == intAnswer {
			correct++
			fmt.Println("That is correct!")
		} else if userIn != intAnswer {
			incorrect++
			fmt.Println("Incorrect Answer")
		}
	}

	return correct, incorrect
}
