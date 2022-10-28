package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// initialise counters
	var questionNumber int

    // Call csvReader and iterate through output
    for _, line := range csvReader("./problems.csv") {
    	questionNumber = questionNumber + 1


    	fmt.Printf("Problem #%d: %s =\n", questionNumber, line[0])
    }
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