package day2

import (
	"regexp"
	"strconv"
	"strings"
)

func isGamePossible(gameData string, maxCubes map[string]int) bool {

	subsets := strings.Split(gameData, "; ")
	for _, subset := range subsets {
		if !subsetPossible(subset, maxCubes) {
			return false
		}
	}
	return true
}

func subsetPossible(subset string, maxCubes map[string]int) bool {
	cubePattern := regexp.MustCompile(`(\d+) (red|green|blue)`)
	matches := cubePattern.FindAllStringSubmatch(subset, -1)

	for _, match := range matches {
		count, _ := strconv.Atoi(match[1])
		color := match[2]

		if count > maxCubes[color] {
			return false
		}
	}
	return true
}

func sumOfPossibleGameIDs(games map[int]string, maxCubes map[string]int) int {
	sum := 0
	for gameID, gameData := range games {
		if isGamePossible(gameData, maxCubes) {
			sum += gameID
		}
	}
	return sum
}

func sumOfPowersOfMinimumSets(games map[int]string) int {
	totalPower := 0
	for _, game := range games {
		minCubes := findMinimumCubes(game)
		totalPower += minCubes["red"] * minCubes["green"] * minCubes["blue"]
	}
	return totalPower
}

func findMinimumCubes(gameData string) map[string]int {
	// Initialize minimum cubes required for each color
	minCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	// Regular expression to find color and number pairs
	colorPattern := regexp.MustCompile(`(\d+) (red|green|blue)`)
	subsets := strings.Split(gameData, "; ")

	for _, subset := range subsets {
		matches := colorPattern.FindAllStringSubmatch(subset, -1)

		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]

			// Update the minimum cubes if the current count is higher
			if count > minCubes[color] {
				minCubes[color] = count
			}
		}
	}

	return minCubes
}
