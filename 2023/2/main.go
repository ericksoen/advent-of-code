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

type Game struct {
	Index int
	Valid bool
	ID    int
	Turns []Turn
}

type Turn struct {
	Index int
	Rolls []Roll
}

type Roll struct {
	Index     int
	Color     string
	Frequency int
}

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

	games := make([]Game, 0)

	for i, line := range split {
		if len(line) == 0 {
			logger.Info("Empty line", "index", i)
			continue
		}

		gameTokens := strings.Split(line, ":")

		if len(gameTokens) != 2 {
			logger.Warn("Malformed line", "index", i, "line", line)
		}

		gameID := gameTokens[0]

		numericGameIdTokens := strings.Split(gameID, " ")

		if len(numericGameIdTokens) != 2 {
			logger.Warn("Invalid Game ID", "line", i, "GameID", numericGameIdTokens)
			continue
		}

		numericGameID, err := strconv.Atoi(numericGameIdTokens[1])

		if err != nil {
			logger.Error("Game ID is not numeric", "lineId", i, "gameId", numericGameIdTokens[1])
			continue
		}

		game := Game{Index: i, ID: numericGameID, Valid: true, Turns: make([]Turn, 0)}

		gamesRaw := strings.Trim(gameTokens[1], " ")

		turnTokens := strings.Split(gamesRaw, ";")

		for j, turnToken := range turnTokens {

			turns := make([]Turn, 0)

			logger.Info("Turn details", "gameID", gameID, "gameIndex", i, "turnIndex", j, "turnToken", turnToken)

			rollToken := strings.Split(turnToken, ",")

			rolls := make([]Roll, 0)
			for k, roll := range rollToken {

				roll = strings.Trim(roll, " ")
				logger.Info("Roll detail", "roll", roll)

				roleDetailTokens := strings.Split(roll, " ")

				if len(roleDetailTokens) != 2 {
					logger.Warn("Roll tokens", "tokens", roleDetailTokens)
				}

				frequency, err := strconv.Atoi(roleDetailTokens[0])

				if err != nil {
					logger.Error("Invalid roll frequency", "tokens", roleDetailTokens)
				}

				color := strings.ToLower(roleDetailTokens[1])

				rolls = append(rolls, Roll{Color: color, Frequency: frequency, Index: k})
				if color == "blue" && frequency > *maxBlue {
					game.Valid = false
					continue
				}

				if color == "green" && frequency > *maxGreen {
					game.Valid = false
					continue
				}

				if color == "red" && frequency > *maxRed {
					game.Valid = false
					continue
				}

			}

			turns = append(turns, Turn{Index: j, Rolls: rolls})
		}

		games = append(games, game)
	}

	logger.Info("All Games", "GameID", games)

	sum := 0

	for _, game := range games {

		if game.Valid {
			sum += game.ID
		}
	}

	logger.Info("", "total", sum)
}
