package main

import (
	"flag"
	"log/slog"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative path to the input file")
var lineDelimiter = flag.String("lineDelimiter", "\n", "The end of line delimiter")

func main() {
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	bytes, err := os.ReadFile(*inputFile)

	if err != nil {
		return
	}

	contents := string(bytes)

	split := strings.Split(contents, *lineDelimiter)
	split = split[:len(split)-1]
	for _, line := range split {
		logger.Info("", "line", line)
	}
}
