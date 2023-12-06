package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Races struct {
	list []Race
}

type Race struct {
	TimeAllowed    int
	DistanceRecord int
}

func ParseData(data string) Races {
	lines := strings.Split(data, "\n")
	timesInLine := extractIntsFromLine(lines[0][6:])
	distancesInLine := extractIntsFromLine(lines[1][10:])
	races := make([]Race, 0)
	for index, timeAllowed := range timesInLine {
		race := Race{TimeAllowed: timeAllowed, DistanceRecord: distancesInLine[index]}
		races = append(races, race)
	}
	return Races{list: races}
}

func ParseDataPart2(data string) Race {
	lines := strings.Split(data, "\n")
	time := extractIntsFromLinePart2(lines[0][6:])
	distance := extractIntsFromLinePart2(lines[1][10:])
	return Race{TimeAllowed: time, DistanceRecord: distance}
}

func (races *Races) GetNumberOfWayToBeatRecord() int {
	result := 1
	for _, race := range races.list {
		result *= race.GetNumberOfWayToBeatRecord()
	}
	return result
}

func (race *Race) GetNumberOfWayToBeatRecord() int {
	distanceToBeat := race.DistanceRecord
	timeAllowed := race.TimeAllowed
	count := 0
	for i := 1; i <= timeAllowed; i++ {
		if IsRecordBeaten(distanceToBeat, timeAllowed, i) {
			count++
		}
	}
	return count
}

func (race *Race) GetNumberOfWayToBeatRecordPart2() int {
	minTimeFloat, maxTimeFloat, error := race.getMinMaxTimeToEqualsRecord()
	if error != nil {
		panic(error)
	}
	fmt.Printf("minTimeFloat: %f, maxTimeFloat: %f\n", minTimeFloat, maxTimeFloat)
	minTime := int(math.Ceil(minTimeFloat))
	maxTime := int(math.Floor(maxTimeFloat))
	return maxTime - minTime + 1
}

func (race *Race) getMinMaxTimeToEqualsRecord() (float64, float64, error) {
	floatTimeAllowed := float64(race.TimeAllowed)
	floatDistanceRecord := float64(race.DistanceRecord)

	// Calculate the discriminant
	discriminant := math.Pow(floatTimeAllowed, 2) - 4*floatDistanceRecord

	if discriminant >= 0 {
		sqrtDiscriminant := math.Sqrt(discriminant)

		// Calculate the two possible solutions
		timeToHoldMax := (floatTimeAllowed + sqrtDiscriminant) / 2
		timeToHoldMin := (floatTimeAllowed - sqrtDiscriminant) / 2

		return timeToHoldMin, timeToHoldMax, nil
	}

	return math.NaN(), math.NaN(), fmt.Errorf("No solution")
}

func IsRecordBeaten(distanceToBeat int, timeAllowed int, timeButtonPushed int) bool {
	time := timeAllowed - timeButtonPushed
	return distanceToBeat < time*timeButtonPushed
}

func extractIntsFromLine(line string) []int {
	list := make([]int, 0)
	for _, val := range strings.Split(line, " ") {
		digit, err := strconv.Atoi(val)
		if err != nil {
		} else {
			list = append(list, digit)
		}
	}
	return list
}

func extractIntsFromLinePart2(line string) int {
	builder := strings.Builder{}
	for _, val := range strings.Split(line, " ") {
		digit, err := strconv.Atoi(val)
		if err != nil {
		} else {
			builder.WriteString(strconv.Itoa(digit))
		}
	}
	digit, _ := strconv.Atoi(builder.String())
	return digit
}
