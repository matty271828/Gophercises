package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
    fmt.Println(csvReader())
}

func csvReader() [][] string{
  // Open the file
  recordFile, err := os.Open("./problems.csv")
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