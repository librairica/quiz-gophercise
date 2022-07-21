package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	fd, error := os.Open("problems.csv")

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Type your answer, then hit enter to see the next question.")
	}

	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}

	numCorrect := 0
	numIncorrect := 0

	reader := bufio.NewReader(os.Stdin)

	for _, element := range records {
		// prints the clue
		fmt.Print(element[0], "\n")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		input = strings.TrimSuffix(input, "\n")

		if input == element[1] {
			numCorrect += 1
		} else {
			numIncorrect += 1
		}
	}
	fmt.Println("Total correct: ", numCorrect, "\n")
	fmt.Println("Total incorrect: ", numIncorrect, "\n")
	fmt.Println("Your final score is: ", numCorrect-numIncorrect, "\n")

}
