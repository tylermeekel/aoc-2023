package main

import (
	"slices"
	"testing"
)

func TestParseGame(t *testing.T) {
	tests := []struct {
		in  string
		out Game
	}{
		{
			in: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			out: Game{
				ID: 1,
				Rounds: []Round{
					{NBlue: 3, NRed: 4},
					{NRed: 1, NGreen: 2, NBlue: 6},
					{NGreen: 2},
				},
			},
		},
	}

	for _, tt := range tests {
		game := ParseGame(tt.in)

		if game.ID != tt.out.ID {
			t.Errorf("Incorrect ID. expected=%d got=%d", tt.out.ID, game.ID)
		}

		if !slices.Equal(game.Rounds, tt.out.Rounds) {
			t.Errorf("Incorrect Rounds. expected=%+v, got=%+v", tt.out.Rounds, game.Rounds)
		}

	}
}

func TestParseRounds(t *testing.T) {
	tests := []struct {
		in  string
		out []Round
	}{
		{
			in: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			out: []Round{
				{NBlue: 3, NRed: 4},
				{NRed: 1, NGreen: 2, NBlue: 6},
				{NGreen: 2},
			},
		},
	}

	for _, tt := range tests {
		rounds := ParseRounds(tt.in)

		if !slices.Equal(rounds, tt.out) {
			t.Errorf("Incorrect rounds: expected=%+v, got=%+v", tt.out, rounds)
		}
	}
}
