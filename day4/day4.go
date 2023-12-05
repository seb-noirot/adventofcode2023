package day4

import (
	"strconv"
	"strings"
)

func calculateSumOfScrathcards(scrathcards []string) int {
	var sum int
	for _, card := range scrathcards {
		totalCard := calculateSumOfScrathcard(card)
		sum += totalCard
	}
	return sum
}

func calculateSumOfScrathcard(card string) int {
	vals := strings.Split(card, ":")
	numbers := strings.Split(vals[1], "|")
	winningNumbers := extractNumber(numbers[0])
	currentNumbers := extractNumber(numbers[1])
	return calculatePoints(winningNumbers, currentNumbers)
}

func calculateSumOfScrathcardNewRules(card string) (int, int) {
	vals := strings.Split(card, ":")
	header := vals[0]
	cardNumberString := strings.Trim(header[4:], " ")
	index, _ := strconv.Atoi(cardNumberString)
	numbers := strings.Split(vals[1], "|")
	winningNumbers := extractNumber(numbers[0])
	currentNumbers := extractNumber(numbers[1])
	return index, calculatePointsNewRules(winningNumbers, currentNumbers)
}

func extractNumber(value string) []int {
	firstSplit := strings.Split(value, " ")
	var numbers []int
	for _, token := range firstSplit {
		val, err := strconv.Atoi(token)
		if err != nil {

		} else {
			numbers = append(numbers, val)
		}
	}
	return numbers
}

func calculatePoints(winningNumbers []int, currentNumbers []int) int {
	points := 0
	for _, currentNumber := range currentNumbers {
		for _, winningNumber := range winningNumbers {
			if winningNumber == currentNumber {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				break
			}
		}
	}
	return points
}

func calculatePointsNewRules(winningNumbers []int, currentNumbers []int) int {
	points := 0
	for _, currentNumber := range currentNumbers {
		for _, winningNumber := range winningNumbers {
			if winningNumber == currentNumber {
				points++
				break
			}
		}
	}
	return points
}

type Cards struct {
	CardsToScratch []int
}

func calculateNumberOfScracthsCard(scrathcards []string) int {
	var counter int
	mapCards := make(map[int]Cards)
	mapCarsAggregates := make(map[int]Cards)
	for _, card := range scrathcards {
		index, cards := calculateCard(card)
		mapCards[index] = cards
		mapCarsAggregates[index] = cards
	}

	for i := 0; i < len(scrathcards); i++ {
		cardIndex := i + 1
		cards := mapCarsAggregates[cardIndex]
		for _, card := range cards.CardsToScratch {
			carsToScratch := append(mapCarsAggregates[card].CardsToScratch, mapCards[card].CardsToScratch...)
			mapCarsAggregates[card] = Cards{CardsToScratch: carsToScratch}
		}
	}

	for i := 1; i <= len(scrathcards); i++ {
		cards, _ := mapCarsAggregates[i]
		counter += len(cards.CardsToScratch) + 1
	}

	return counter
}

func calculateCard(card string) (int, Cards) {
	index, sum := calculateSumOfScrathcardNewRules(card)
	var cards []int
	if sum > 0 {
		for indexCard := index + 1; indexCard <= index+sum; indexCard++ {
			cards = append(cards, indexCard)
		}
	}
	return index, Cards{CardsToScratch: cards}
}
