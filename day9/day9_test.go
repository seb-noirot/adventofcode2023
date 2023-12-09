package day9

import (
	"adventofcode/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var demoInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestParseData(t *testing.T) {
	got := ParseData(demoInput)
	expect := []History{{Numbers: []int{0, 3, 6, 9, 12, 15}}, {Numbers: []int{1, 3, 6, 10, 15, 21}}, {Numbers: []int{10, 13, 16, 21, 30, 45}}}

	assert.Equal(t, expect, got)
}

func TestExtrapolate(t *testing.T) {
	history := History{Numbers: []int{0, 3, 6, 9, 12, 15}}
	got := history.Extrapolate()
	expect := Extrapolation{
		Histories: []History{
			{Numbers: []int{0, 3, 6, 9, 12, 15}},
			{Numbers: []int{3, 3, 3, 3, 3}},
			{Numbers: []int{0, 0, 0, 0}},
		},
	}

	assert.Equal(t, expect, got)
}

func TestCalculateNextStep(t *testing.T) {
	extrapolation := Extrapolation{
		Histories: []History{
			{Numbers: []int{0, 3, 6, 9, 12, 15}},
			{Numbers: []int{3, 3, 3, 3, 3}},
			{Numbers: []int{0, 0, 0, 0}},
		},
	}
	nextStep := extrapolation.CalculateNextStep()
	expect := Extrapolation{
		Histories: []History{
			{Numbers: []int{0, 3, 6, 9, 12, 15, 18}},
			{Numbers: []int{3, 3, 3, 3, 3, 3}},
			{Numbers: []int{0, 0, 0, 0, 0}},
		},
	}

	assert.Equal(t, expect, nextStep)
}

func TestGetLatestHistoryValue(t *testing.T) {
	extrapolation := Extrapolation{
		Histories: []History{
			{Numbers: []int{0, 3, 6, 9, 12, 15, 18}},
			{Numbers: []int{3, 3, 3, 3, 3, 3}},
			{Numbers: []int{0, 0, 0, 0, 0}},
		},
	}
	got := extrapolation.GetLatestHistoryValue()
	expect := 18

	assert.Equal(t, expect, got)
}

func TestGetLatestHistoryValue2(t *testing.T) {
	extrapolation := Extrapolation{
		Histories: []History{
			{Numbers: []int{10, 13, 16, 21, 30, 45, 68}},
			{Numbers: []int{3, 3, 5, 9, 15, 23}},
			{Numbers: []int{0, 2, 4, 6, 8}},
			{Numbers: []int{2, 2, 2, 2}},
			{Numbers: []int{0, 0, 0}},
		},
	}
	got := extrapolation.GetLatestHistoryValue()
	expect := 68

	assert.Equal(t, expect, got)
}

func TestGetNextHistoryDigit(t *testing.T) {
	history := History{Numbers: []int{10, 13, 16, 21, 30, 45}}
	latestDigit := history.GetNextHistoryDigit()
	expectedDigit := 68
	assert.Equal(t, expectedDigit, latestDigit)
}

func TestGetNextHistoryDigitSum(t *testing.T) {
	latestDigit := GetNextHistoryDigitSum(demoInput)
	expectedDigit := 114
	assert.Equal(t, expectedDigit, latestDigit)
}

func TestGetNextHistoryDigitSumPart1(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)
	latestDigit := GetNextHistoryDigitSum(content)
	expectedDigit := 1992273652
	assert.Equal(t, expectedDigit, latestDigit)
}

func TestGetPreviousHistoryDigitSum(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)
	latestDigit := GetPreviousHistoryDigitSum(content)
	expectedDigit := 1012
	assert.Equal(t, expectedDigit, latestDigit)
}

func TestGetPreviousHistoryDigitSumDemo(t *testing.T) {
	history := History{Numbers: []int{10, 13, 16, 21, 30, 45}}
	latestDigit := history.GetPreviousHistoryDigit()
	expectedDigit := 5
	assert.Equal(t, expectedDigit, latestDigit)
}

func TestNegativeData(t *testing.T) {
	history := History{Numbers: []int{3, -2, -7, -12, -17, -22, -27, -32, -37, -42, -47, -52, -57, -62, -67, -72, -77, -82, -87, -92, -97}}
	latestDigit := history.GetNextHistoryDigit()
	expectedDigit := -102
	assert.Equal(t, expectedDigit, latestDigit)
}
