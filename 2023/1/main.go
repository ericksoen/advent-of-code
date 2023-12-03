package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative path to the input file")
var lineDelimiter = flag.String("lineDelimiter", "\n", "The end of line delimiter")
var useWordReplacement = flag.Bool("wordReplacement", true, "Determines if word tokens like `eight` and `one` should be replaced with their numeric equivalent")
var maxLine = flag.Int("maxLine", -1, "The maximum line index to process. -1 defaults to all.")
var minLine = flag.Int("minLine", 0, "The minimum line index to process.")

var numericSearchItems = []SearchItem{
	{"1", 1},
	{"2", 2},
	{"3", 3},
	{"4", 4},
	{"5", 5},
	{"6", 6},
	{"7", 7},
	{"8", 8},
	{"9", 9},
}

var alphaSearchItems = []SearchItem{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

type SearchItem struct {
	Token string
	Value int
}

type MatchItem struct {
	Index int
	Value int
}

func finder(line string, characters []SearchItem) []int {

	var matchItems []MatchItem
	for _, token := range characters {

		tokenCount := strings.Count(line, token.Token)

		if tokenCount == 0 {
			continue
		}

		searchPos := 0
		for i := 0; i < tokenCount; i++ {

			matchPos := strings.Index(line[searchPos:], token.Token)

			if matchPos == -1 {
				log.Printf("Should have had a match at this point")
			}

			matchItems = append(matchItems, MatchItem{matchPos + searchPos, token.Value})

			searchPos += matchPos + 1
		}
	}

	sort.Slice(matchItems, func(i, j int) bool {
		return matchItems[i].Index < matchItems[j].Index
	})

	lineTokens := make([]int, len(matchItems))
	for indexPos, val := range matchItems {
		lineTokens[indexPos] = val.Value
	}

	return lineTokens
}

func main() {
	flag.Parse()

	log.Printf("The input file name = %s", *inputFile)
	log.Printf("Using wordReplacement value = %t", *useWordReplacement)
	fileBytes, err := os.ReadFile(*inputFile)

	if err != nil {
		log.Fatalf("Encountered error = %s while reading file", err)
	}
	contents := string(fileBytes)

	split := strings.Split(contents, *lineDelimiter)

	aggregateTotal := 0

	searchItems := numericSearchItems

	if *useWordReplacement {
		searchItems = append(searchItems, alphaSearchItems...)
	}

	log.Printf("Using search items = %+v", searchItems)

	totalLines := len(split)
	if *maxLine != -1 {
		totalLines = *maxLine
	}
	for i, line := range split[*minLine:totalLines] {

		if len(line) == 0 {
			log.Printf("Skipping empty line with index = %d", i)
			continue
		}
		lineTokens := finder(line, searchItems)

		firstToken := lineTokens[0]
		lastToken := lineTokens[len(lineTokens)-1]

		lineTotal, err := strconv.Atoi(fmt.Sprintf("%d%d", firstToken, lastToken))

		if err != nil {
			panic(err)
		}

		log.Printf("[Line %d] Total = %d for token = %s.", i, lineTotal, line)
		aggregateTotal += lineTotal

		if i%10 == 0 {
			log.Printf("[Line %d] Aggregate total line = %d.", i, aggregateTotal)
		}
	}
	log.Printf("The file totals = %d", aggregateTotal)
}
