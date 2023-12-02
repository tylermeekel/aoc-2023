package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	total := 0
	expectedTotal := 281

	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	for _, line := range lines {
		total += GetDigits(line)
	}

	if total != expectedTotal{
		t.Errorf("Total incorrect: expected=%d got=%d", expectedTotal, total)
	}
}

func TestGetDigits(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tt := range tests {
		dig := GetDigits(tt.in)
		if dig != tt.out {
			t.Errorf("incorrect output, in=%s expected=%d, got=%d",tt.in, tt.out, dig)
		}
	}
}
