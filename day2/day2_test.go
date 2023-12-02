package day2

import (
	"adventofcode/helper"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGamePossible(t *testing.T) {

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	testCases := []struct {
		gameID     int
		gameData   string
		isPossible bool
	}{
		{1, "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{2, "1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{3, "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{4, "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{5, "6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	for _, tc := range testCases {
		possible := isGamePossible(tc.gameData, maxCubes)
		assert.Equal(t, tc.isPossible, possible)
	}
}

func TestSumOfPowersOfMinimumSets(t *testing.T) {
	games := map[int]string{
		1: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		2: "1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		3: "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		4: "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		5: "6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	expectedSum := 2286 // Adjust based on your full input data
	assert.Equal(t, expectedSum, sumOfPowersOfMinimumSets(games))
}

func TestSumOfPowersOfMinimumSetsWithInput(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	games, err := processInput(input)
	if err != nil {
		panic(err)
	}
	sum := sumOfPowersOfMinimumSets(games)
	expectedSum := 67953
	assert.Equal(t, expectedSum, sum)
}

func TestSumOfPossibleGameIDs(t *testing.T) {

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := map[int]string{
		1: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		2: "1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		3: "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		4: "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		5: "6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	expectedSum := 8 // Sum of IDs of games 1, 2, and 5
	assert.Equal(t, expectedSum, sumOfPossibleGameIDs(games, maxCubes))
}

func TestCalculateSum(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	games, err := processInput(input)
	if err != nil {
		panic(err)
	}
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := sumOfPossibleGameIDs(games, maxCubes)
	expectedSum := 2278
	assert.Equal(t, expectedSum, sum)
}

func processInput(lines []string) (map[int]string, error) {

	games := make(map[int]string)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue // Skip malformed lines
		}

		gameIDStr := strings.TrimSpace(strings.TrimPrefix(parts[0], "Game"))
		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			return nil, err // Handle parse error
		}

		gameData := strings.TrimSpace(parts[1])
		games[gameID] = gameData
	}
	return games, nil
}
