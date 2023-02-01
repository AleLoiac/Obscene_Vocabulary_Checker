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

func censor(word string) {
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func checkWord(file *os.File) {

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var obscene bool

	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == strings.ToLower(word) {
			obscene = true
			break
		} else {
			obscene = false
		}
	}
	if obscene == true {
		censor(word)
	} else {
		fmt.Println(word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	_, err := fmt.Scan(&fileName)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for {

		file := readFile(fileName)
		defer file.Close() //defer the closure

		_, err2 := fmt.Scan(&word)
		if err2 != nil {
			log.Fatalf("error: %v", err2)
		}
		if strings.ToLower(word) == "exit" {
			break
		}
		checkWord(file)
	}
}
