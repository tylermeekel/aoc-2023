package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID     int
	Rounds []Round
}

type Round struct {
	NRed   int
	NGreen int
	NBlue  int
}

func main() {
	totalPower := 0

	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error opening file:", err.Error())
	}

	var games []Game
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		games = append(games, ParseGame(reader.Text()))
	}

	for _, game := range games{
		r, g, b := GetMinCounts(game)
		totalPower += r*g*b
	}

	fmt.Println("Total power is:", totalPower)
}

func GetMinCounts(game Game) (r, g, b int) {
	for _, round := range game.Rounds{
		r = max(r, round.NRed)
		g = max(g, round.NGreen)
		b = max(b, round.NBlue)
	}

	return
}

func ParseGame(in string) Game {
	var out Game

	splitStr := strings.Split(in, ": ")
	gameIDStr := strings.Split(splitStr[0], " ")[1]

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		fmt.Println("Error parsing ID")
	}

	out.ID = gameID

	out.Rounds = ParseRounds(splitStr[1])

	return out
}

func ParseRounds(in string) []Round {
	var rounds []Round
	roundsStr := strings.Split(in, "; ")

	for _, roundStr := range roundsStr {
		var round Round

		countsStr := strings.Split(roundStr, ", ")
		for _, countStr := range countsStr {
			split := strings.Split(countStr, " ")
			count, err := strconv.Atoi(split[0])
			if err != nil {
				fmt.Println("Error parsing count:", err.Error())
			}
			color := split[1]

			switch color {
			case "red":
				round.NRed = count
			case "green":
				round.NGreen = count
			case "blue":
				round.NBlue = count
			}
		}

		rounds = append(rounds, round)
	}

	return rounds
}
