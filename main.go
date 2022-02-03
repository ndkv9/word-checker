package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// take the file name from user's input
	var fileName string
	fmt.Print("> ")
	fmt.Scan(&fileName)

	// handle open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// read the file word by word
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var check string
	wordList := make(map[string]struct{})

	// save each word into a set
	for scanner.Scan() {
		wordList[scanner.Text()] = struct{}{}
	}

	fmt.Print("> ")
	fmt.Scan(&check)

	if _, ok := wordList[strings.ToLower(check)]; ok {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// handle err during scanning the file
	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

}
