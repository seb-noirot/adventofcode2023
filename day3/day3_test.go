package day3

import (
	"adventofcode/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfPartNumbers(t *testing.T) {
	testCases := []struct {
		schematic   []string
		expectedSum int
	}{
		{
			schematic: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				".......#..",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expectedSum: 4361,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedSum, sumOfPartNumber(tc.schematic))
	}
}

func TestSumOfPartGearsRatio(t *testing.T) {
	testCases := []struct {
		schematic   []string
		expectedSum int
	}{
		{
			schematic: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				".......#..",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expectedSum: 467835,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedSum, sumOfPartGearsRatio(tc.schematic))
	}
}

func TestPart1(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	sum := sumOfPartNumber(input)
	expectedSum := 540131
	assert.Equal(t, expectedSum, sum)
}

func TestPart2(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	sum := sumOfPartGearsRatio(input)
	expectedSum := 86879020
	assert.Equal(t, expectedSum, sum)
}
