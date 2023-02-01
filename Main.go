package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string
var sentence string

func printWords(file *os.File) { //prints all names in the file on a different row

	scanner := bufio.NewScanner(file)

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
	fmt.Print(" ")
}

func checkWord(file *os.File) {

	scanner := bufio.NewScanner(file) // creates a new Scanner for the file
	scanner.Split(bufio.ScanWords)

	sentenceScanner := bufio.NewScanner(strings.NewReader(sentence))
	sentenceScanner.Split(bufio.ScanWords)

	var obscene bool

	for sentenceScanner.Scan() {
		for scanner.Scan() {
			if strings.ToLower(scanner.Text()) == strings.ToLower(sentenceScanner.Text()) {
				obscene = true
				break
			} else {
				obscene = false
			}
		}
		if obscene == true {
			censor(sentenceScanner.Text())
		} else {
			fmt.Print(sentenceScanner.Text() + " ")
		}
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

	file := readFile(fileName)
	defer file.Close() //defer the closure

	for {
		_, err2 := fmt.Scan(&sentence)
		if err2 != nil {
			log.Fatalf("error: %v", err2)
		}
		if strings.ToLower(sentence) == "exit" {
			fmt.Println("Bye!")
			break
		}
		checkWord(file)

		// reset the position of the file back to the start after each iteration of the loop,
		// so that the bufio.Scanner reads the contents of the file from the beginning each time.
		_, err3 := file.Seek(0, 0)
		if err3 != nil {
			return
		}
	}
}
