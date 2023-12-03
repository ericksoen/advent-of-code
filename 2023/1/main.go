package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative path to the input file")
var lineDelimiter = flag.String("lineDelimiter", "\n", "The end of line delimiter")

var nonNumericRegexp = regexp.MustCompile(`[^0-9]+`)

func finder(line string) []int {

	sanitized := nonNumericRegexp.ReplaceAllString(line, "")

	lineTokens := make([]int, len(sanitized))
	for indexPos, val := range sanitized {
		i, err := strconv.Atoi(string(val))

		if err != nil {
			panic(err)
		}

		lineTokens[indexPos] = i
	}

	return lineTokens
}

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
	for i, line := range split {

		lineTokens := finder(line)

		firstToken := lineTokens[0]
		lastToken := lineTokens[len(lineTokens)-1]

		lineTotal, err := strconv.Atoi(fmt.Sprintf("%d%d", firstToken, lastToken))

		if err != nil {
			panic(err)
		}

		aggregateTotal += lineTotal
		log.Printf("[Line %d]: Total = %d for token = %s", i, lineTotal, line)
	}
	log.Printf("The file totals = %d", aggregateTotal)
}
