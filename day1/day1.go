package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getCalibrationSum(input []string) int {
	var sum int
	for _, s := range input {
		sum += getCalibrationValue(s)
	}
	fmt.Printf("sum: %d\n", sum)
	return sum
}

func getCalibrationValue(s string) int {
	firstDigit, lastDigit := findDigits(s)
	value, _ := strconv.Atoi(firstDigit + lastDigit)
	fmt.Printf(", value: %d\n", value)
	return value
}

func findLast(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsLetter(r) {
			word := string(runes[i:])
			if digit, ok := keyIncludeInDigitMap(word); ok {
				return digit
			}
		} else if unicode.IsDigit(r) {
			return string(r)
		}
	}

	return ""
}

func findFirst(s string) string {
	var currentWord strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			currentWord.WriteRune(r)
			if digit, ok := keyIncludeInDigitMap(currentWord.String()); ok {
				return digit
			}
		} else if unicode.IsDigit(r) {
			return string(r)
		}
	}
	return ""
}

func findDigits(s string) (string, string) {
	firstDigit := findFirst(s)
	lastDigit := findLast(s)
	return firstDigit, lastDigit
}

func keyIncludeInDigitMap(s string) (string, bool) {
	for key := range digitMap {
		if strings.Contains(s, key) {
			return digitMap[key], true
		}
	}
	return "", false
}
