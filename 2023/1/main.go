package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative path to the input file")
var lineDelimiter = flag.String("lineDelimiter", "\n", "The end of line delimiter")

func main() {
	flag.Parse()

	fileBytes, err := os.ReadFile(*inputFile)

	if err != nil {
		return
	}

	contents := string(fileBytes)

	split := strings.Split(contents, *lineDelimiter)
	split = split[:len(split)-1]

	aggregateTotal := 0
	for _, line := range split {
		// log.Printf("The line = %s", line)
		var lineTokens []int
		for _, char := range line {

			i, err := strconv.Atoi(string(char))

			if err != nil {
				continue
			}

			lineTokens = append(lineTokens, i)
		}

		firstToken := lineTokens[0]
		lastToken := lineTokens[len(lineTokens)-1]

		lineTotal, err := strconv.Atoi(fmt.Sprintf("%d%d", firstToken, lastToken))
		// log.Printf("The line total = %d", lineTotal)

		if err != nil {
			panic(err)
		}

		aggregateTotal += lineTotal

		log.Printf("The file totals = %d", aggregateTotal)
	}
}
