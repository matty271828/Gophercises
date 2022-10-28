package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// initialise counters
	var questionNumber int
	var score int

    // Call csvReader and iterate through output
    for _, line := range csvReader("./problems.csv") {
    	questionNumber = questionNumber + 1

    	// Ask user the question
    	fmt.Printf("Problem #%d: %s = ", questionNumber, line[0])

    	// Get user input
    	var answer string
    	fmt.Scan(&answer)

    	// Compare user's answer to solution
    	if answer == line[1] {
    		score = score + 1
    	}

    }

    // Output user score and number of questions
    fmt.Printf("You have correctly answered %d of %d questions\n", score, questionNumber)
}

func csvReader(fileName string) [][] string{
  // Open the file
  recordFile, err := os.Open(fileName)
  if err != nil {
   fmt.Println("An error encountered ::", err)
  }
  // Initialize the reader
  reader := csv.NewReader(recordFile)

  // Read all the records
  records, _ := reader.ReadAll()

  // Copy records onto 2d slice
  var questionsAndAnswers [][] string = records

  // Return the slice
  return questionsAndAnswers
}