package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative path to the input file")
var lineDelimiter = flag.String("lineDelimiter", "\n", "The end of line delimiter")

func main() {
	flag.Parse()

	bytes, err := os.ReadFile(*inputFile)

	if err != nil {
		return
	}

	contents := string(bytes)

	split := strings.Split(contents, *lineDelimiter)
	split = split[:len(split)-1]
	for _, line := range split {
		log.Printf("The line contents are %s", line)
	}
}
