package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string
var word string

func printWords(file *os.File) { //prints all names in the file on a different row

	scanner := bufio.NewScanner(file) // create a new Scanner for the file

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func checkWord(file *os.File) {

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	x := ""
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == strings.ToLower(word) {
			x = "True"
			break
		} else {
			x = "False"
		}
	}
	fmt.Println(x)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	_, err := fmt.Scan(&fileName, &word)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file := readFile(fileName)
	defer file.Close() //defer the closure

	checkWord(file)
}
