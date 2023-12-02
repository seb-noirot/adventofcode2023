package day1

import (
	"adventofcode/helper"
	_ "adventofcode/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateSum(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	assert.NoError(t, err)

	sum := getCalibrationSum(input)
	expectedSum := 53592
	assert.Equal(t, expectedSum, sum)
}

func TestGetCalibrationValue(t *testing.T) {
	assert.Equal(t, 12, getCalibrationValue("1abc2"))
	assert.Equal(t, 38, getCalibrationValue("pqr3stu8vwx"))
	assert.Equal(t, 15, getCalibrationValue("a1b2c3d4e5f"))
	assert.Equal(t, 77, getCalibrationValue("treb7uchet"))
}

func TestGetCalibrationSum(t *testing.T) {
	assert.Equal(t, 142, getCalibrationSum([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}))
}

func TestGetCalibrationValueWithNewMethod1(t *testing.T) {
	got := getCalibrationValue("eightwothree")
	assert.Equal(t, 83, got)
}

func TestGetCalibrationValueNewMethod(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"mxmkjvgsdzfhseightonetwoeight7", 87},
		{"jjhxddmg5mqxqbgfivextlcpnvtwothreetwonerzk", 52},
		{"bm6fourghmnrnsmtwotwofournssrseven", 67},
		{"eightwo", 82},
	}

	for _, tc := range testCases {
		got := getCalibrationValue(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func TestGetSumValueNewMethod(t *testing.T) {
	assert.Equal(t, 281, getCalibrationSum([]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}))
}
