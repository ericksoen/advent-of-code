package main

import (
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative path to the input file")
var lineDelimiter = flag.String("lineDelimiter", "\n", "The end of line delimiter")
var maxGreen = flag.Int("maxGreen", -1, "The max number of greens")
var maxBlue = flag.Int("maxBlue", -1, "The max number of bluess")
var maxRed = flag.Int("maxRed", -1, "The max number of reds")

func main() {
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	if *maxGreen == -1 || *maxBlue == -1 || *maxRed == -1 {
		logger.Error("Invalid inputs", "green", *maxGreen, "blue", *maxBlue, "red", *maxBlue)
		os.Exit(-1)
	}

	bytes, err := os.ReadFile(*inputFile)

	if err != nil {
		return
	}

	contents := string(bytes)

	split := strings.Split(contents, *lineDelimiter)

	validGameIDs := make([]int, 0)

	for i, line := range split {
		if len(line) == 0 {
			logger.Info("Empty line", "index", i)
			continue
		}

		isValidGame := true

		tokens := strings.Split(line, ":")

		if len(tokens) != 2 {
			logger.Warn("Malformed line", "index", i, "line", line)
		}

		gameID := tokens[0]
		games := strings.Trim(tokens[1], " ")

		gameTokens := strings.Split(games, ";")

		for i, gameToken := range gameTokens {
			logger.Info("Game details", "gameID", gameID, "gameIndex", i, "gameToken", gameToken)

			rollDetails := strings.Split(gameToken, ",")

			for _, rollDetail := range rollDetails {

				rollDetail = strings.Trim(rollDetail, " ")
				logger.Info("Roll detail", "detail", rollDetail)

				rollTokens := strings.Split(rollDetail, " ")

				if len(rollTokens) != 2 {
					logger.Warn("Roll tokens", "tokens", rollTokens)
				}

				frequency, err := strconv.Atoi(rollTokens[0])

				if err != nil {
					logger.Error("Invalid roll frequency", "tokens", rollTokens)
				}

				color := strings.ToLower(rollTokens[1])

				if color == "blue" && frequency > *maxBlue {
					isValidGame = false
					continue
				}

				if color == "green" && frequency > *maxGreen {
					isValidGame = false
					continue
				}

				if color == "red" && frequency > *maxRed {
					isValidGame = false
					continue
				}

			}
		}

		if isValidGame {
			numericGameIdTokens := strings.Split(gameID, " ")

			if len(numericGameIdTokens) != 2 {
				logger.Warn("Invalid Game ID", "GameID", numericGameIdTokens)
				continue
			}

			numericGameID, err := strconv.Atoi(numericGameIdTokens[1])

			if err != nil {
				logger.Error("Game ID is not numeric", "gameId", numericGameIdTokens[1])
				continue
			}

			validGameIDs = append(validGameIDs, numericGameID)

		}
	}

	logger.Info("Valid Game IDs", "GameID", validGameIDs)

	sum := 0

	for _, validGameId := range validGameIDs {
		sum += validGameId
	}

	logger.Info("", "total", sum)
}
