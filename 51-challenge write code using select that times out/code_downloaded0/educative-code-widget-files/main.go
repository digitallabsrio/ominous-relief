package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// struct
type problem struct {
	q, a string
}

// The qa.csv is a sample file, you can create your own quiz

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}


// function read csv file
func readCSVFile(filename string) [][]string {
	
}

// function to store result into problem array
func parseProblems(records [][]string) []problem {

}

func main() {
    // timer 
	timer := time.NewTimer(time.Duration(5) * time.Second)
	correct, incorrect := 0, 0
	// main logic to timesout after 30 seconds.

}