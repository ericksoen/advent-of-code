package main

import (
	"flag"
	"fmt"
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
	Index    int
	Valid    bool
	ID       int
	Turns    []Turn
	MaxGreen int
	MaxBlue  int
	MaxRed   int
}

func NewGame(index, id, maxGreen, maxBlue, maxRed int) *Game {

	return &Game{
		Index:    index,
		ID:       id,
		Valid:    true,
		MaxGreen: maxGreen,
		MaxBlue:  maxBlue,
		MaxRed:   maxRed,
		Turns:    make([]Turn, 0),
	}
}

func (g *Game) AddTurn(t Turn) {
	g.Turns = append(g.Turns, t)

	if t.Blue > g.MaxBlue {
		g.Valid = false
	}

	if t.Green > g.MaxGreen {
		g.Valid = false
	}

	if t.Red > g.MaxRed {
		g.Valid = false
	}
}

type Turn struct {
	Index int
	Blue  int
	Green int
	Red   int
}

func main() {
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
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

		game := NewGame(i, numericGameID, *maxGreen, *maxBlue, *maxRed)

		gamesRaw := strings.Trim(gameTokens[1], " ")

		turnTokens := strings.Split(gamesRaw, ";")

		for j, turnToken := range turnTokens {

			logger.Info("Turn details", "gameID", gameID, "gameIndex", i, "turnIndex", j, "turnToken", turnToken)

			rollToken := strings.Split(turnToken, ",")

			red := 0
			blue := 0
			green := 0
			for _, roll := range rollToken {

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

				if color == "blue" {
					blue = frequency
				}

				if color == "green" {
					green = frequency
				}

				if color == "red" {
					red = frequency
				}
			}

			game.AddTurn(Turn{Index: j, Blue: blue, Red: red, Green: green})

			logger.Info("Turn counter", "Id", game.ID, "turns", len(game.Turns))
		}

		games = append(games, *game)
	}

	sum := 0

	for _, game := range games {
		if game.Valid {
			logger.Warn("Game status", "id", game.ID, "status", game.Valid)
			sum += game.ID
		}
	}

	fmt.Printf("The total = %d", sum)
}
