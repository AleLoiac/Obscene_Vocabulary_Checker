package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var fileName string

func main() {

	_, err := fmt.Scan(&fileName)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // create a new Scanner for the file

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text()) // the Text() function converts the scanned bytes to a string
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
