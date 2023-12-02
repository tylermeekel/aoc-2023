package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Match struct {
	idx     int
	content string
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error opening file:", err.Error())
	}

	answer := 0
	lines := 0

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		answer += GetDigits(reader.Text())
		lines++
	}

	fmt.Println("Answer is:", answer)
}

func GetDigits(in string) int {
	digitMap := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	matches := []Match{}

	for k := range digitMap {
		first := strings.Index(in, k)
		if first != -1 {
			matches = append(matches, Match{idx: first, content: k})
		}

		last := strings.LastIndex(in, k)
		// prevent repeated matches if only one instance found
		if last != -1 && last != first {
			matches = append(matches, Match{idx: last, content: k})
		}
	}

	slices.SortFunc(matches, sortMatches)

	//find first match and put associated int in 10s spot, put last match in ones spot
	ans := digitMap[matches[0].content]*10 + digitMap[matches[len(matches)-1].content]

	return ans
}

// Sorts matches in ascending index order
func sortMatches(a, b Match) int {
	return a.idx - b.idx
}
