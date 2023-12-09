package day9

import (
	"strconv"
	"strings"
)

type History struct {
	Numbers []int
}

type Extrapolation struct {
	Histories []History
}

func ParseData(data string) []History {
	lines := strings.Split(data, "\n")
	histories := make([]History, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		history := ParseLine(line)
		histories = append(histories, history)
	}
	return histories
}

func ParseLine(line string) History {
	history := History{}
	for _, val := range strings.Split(line, " ") {
		if strings.Trim(val, " ") == "" {
			continue
		}
		intVal, _ := strconv.Atoi(val)
		history.Numbers = append(history.Numbers, intVal)
	}
	return history
}

func (history History) Last() bool {
	for _, val := range history.Numbers {
		if val != 0 {
			return false
		}
	}
	return true
}

func (extrapolation Extrapolation) Finished() bool {
	for _, history := range extrapolation.Histories {
		if history.Last() {
			return true
		}
	}
	return false
}

func (history History) calculateNext() History {
	newHistory := History{Numbers: make([]int, 0)}
	for index, val := range history.Numbers[1:] {
		newHistory.Numbers = append(newHistory.Numbers, val-history.Numbers[index])
	}
	return newHistory
}

func (extrapolation Extrapolation) calculateNext() Extrapolation {
	lastHistory := extrapolation.Histories[len(extrapolation.Histories)-1]
	newHistory := lastHistory.calculateNext()
	extrapolation.Histories = append(extrapolation.Histories, newHistory)
	return extrapolation
}

func (history History) Extrapolate() Extrapolation {
	extrapolation := Extrapolation{Histories: []History{history}}
	for !extrapolation.Finished() {
		extrapolation = extrapolation.calculateNext()
	}
	return extrapolation
}

func (history *History) calculateNextStep(previousHistory History) {
	newVal := history.Numbers[len(history.Numbers)-1] + previousHistory.Numbers[len(previousHistory.Numbers)-1]
	history.Numbers = append(history.Numbers, newVal)
}

func (history *History) calculatePreviousStep(previousHistory History) {
	newVal := history.Numbers[0] - previousHistory.Numbers[0]
	history.Numbers = append([]int{newVal}, history.Numbers...)
}

func (extrapolation Extrapolation) CalculateNextStep() Extrapolation {
	lastHistory := extrapolation.Histories[len(extrapolation.Histories)-1]
	lastHistory.Numbers = append(lastHistory.Numbers, 0)
	newHistories := make([]History, len(extrapolation.Histories))
	newHistories[len(newHistories)-1] = lastHistory
	for i := len(extrapolation.Histories) - 1; i > 0; i-- {
		newHistories[i-1] = extrapolation.Histories[i-1]
		newHistories[i-1].calculateNextStep(newHistories[i])
	}

	return Extrapolation{Histories: newHistories}
}

func (extrapolation Extrapolation) CalculatePreviousStep() Extrapolation {
	lastHistory := extrapolation.Histories[len(extrapolation.Histories)-1]
	lastHistory.Numbers = append([]int{0}, lastHistory.Numbers...)
	newHistories := make([]History, len(extrapolation.Histories))
	newHistories[len(newHistories)-1] = lastHistory
	for i := len(extrapolation.Histories) - 1; i > 0; i-- {
		newHistories[i-1] = extrapolation.Histories[i-1]
		newHistories[i-1].calculatePreviousStep(newHistories[i])
	}

	return Extrapolation{Histories: newHistories}
}

func (extrapolation Extrapolation) GetLatestHistoryValue() int {
	numbers := extrapolation.Histories[0].Numbers
	return numbers[len(numbers)-1]
}

func (extrapolation Extrapolation) GetFirstHistoryValue() int {
	numbers := extrapolation.Histories[0].Numbers
	return numbers[0]
}

func (history History) GetNextHistoryDigit() int {
	extrapolation := history.Extrapolate()
	extrapolation = extrapolation.CalculateNextStep()
	return extrapolation.GetLatestHistoryValue()
}

func (history History) GetPreviousHistoryDigit() int {
	extrapolation := history.Extrapolate()
	extrapolation = extrapolation.CalculatePreviousStep()
	return extrapolation.GetFirstHistoryValue()
}

func GetNextHistoryDigitSum(data string) int {
	histories := ParseData(data)
	sum := 0
	for _, history := range histories {
		sum += history.GetNextHistoryDigit()
	}
	return sum
}

func GetPreviousHistoryDigitSum(data string) int {
	histories := ParseData(data)
	sum := 0
	for _, history := range histories {
		sum += history.GetPreviousHistoryDigit()
	}
	return sum
}
