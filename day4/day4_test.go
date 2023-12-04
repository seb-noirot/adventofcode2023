package day4

import (
	"adventofcode/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSumOfScrathcards(t *testing.T) {
	scrathcards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	sum := calculateSumOfScrathcards(scrathcards)
	expectedSum := 13
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcards_1(t *testing.T) {
	sum := calculateSumOfScrathcards([]string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"})
	expectedSum := 8
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcards_2(t *testing.T) {
	sum := calculateSumOfScrathcards([]string{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"})
	expectedSum := 2
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcards_3(t *testing.T) {
	sum := calculateSumOfScrathcards([]string{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"})
	expectedSum := 2
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcards_4(t *testing.T) {
	sum := calculateSumOfScrathcards([]string{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"})
	expectedSum := 1
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcards_5(t *testing.T) {
	sum := calculateSumOfScrathcards([]string{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"})
	expectedSum := 0
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcards_6(t *testing.T) {
	sum := calculateSumOfScrathcards([]string{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"})
	expectedSum := 0
	assert.Equal(t, expectedSum, sum)
}

func TestCalculateSumOfScrathcardsNewRules(t *testing.T) {
	scrathcards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	total := calculateNumberOfScracthsCard(scrathcards)
	expectedTotal := 30
	assert.Equal(t, expectedTotal, total)
}

func TestPart1(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	sum := calculateSumOfScrathcards(input)
	expectedSum := 17803
	assert.Equal(t, expectedSum, sum)
}

func TestPart2(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	total := calculateNumberOfScracthsCard(input)
	expectedTotal := 5554894
	assert.Equal(t, expectedTotal, total)
}
