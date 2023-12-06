package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var inputExercise = `Time:        58     81     96     76
Distance:   434   1041   2219   1218`

var inputExample = `Time:      7  15   30
Distance:  9  40  200`

func TestParseData(t *testing.T) {

	expected := Races{list: []Race{{TimeAllowed: 7, DistanceRecord: 9}, {TimeAllowed: 15, DistanceRecord: 40}, {TimeAllowed: 30, DistanceRecord: 200}}}

	got := ParseData(inputExample)
	assert.Equal(t, expected, got)
}

func TestGetNumberOfWayToBeatRecord_1(t *testing.T) {
	races := Races{list: []Race{{TimeAllowed: 7, DistanceRecord: 9}}}
	expected := 4

	got := races.GetNumberOfWayToBeatRecord()
	assert.Equal(t, expected, got)
}

func TestGetNumberOfWayToBeatRecord_2(t *testing.T) {
	races := Races{list: []Race{{TimeAllowed: 15, DistanceRecord: 40}}}
	expected := 8

	got := races.GetNumberOfWayToBeatRecord()
	assert.Equal(t, expected, got)
}

func TestGetNumberOfWayToBeatRecord_3(t *testing.T) {
	races := Races{list: []Race{{TimeAllowed: 30, DistanceRecord: 200}}}
	expected := 9

	got := races.GetNumberOfWayToBeatRecord()
	assert.Equal(t, expected, got)
}

func TestGetNumberOfWayToBeatRecord(t *testing.T) {
	races := Races{list: []Race{{TimeAllowed: 7, DistanceRecord: 9}, {TimeAllowed: 15, DistanceRecord: 40}, {TimeAllowed: 30, DistanceRecord: 200}}}
	expected := 288

	got := races.GetNumberOfWayToBeatRecord()
	assert.Equal(t, expected, got)
}

func TestGetNumberOfWayToBeatRecord_Part1(t *testing.T) {
	races := ParseData(inputExercise)
	expected := 1159152

	got := races.GetNumberOfWayToBeatRecord()
	assert.Equal(t, expected, got)
}

func TestParseDataPart2(t *testing.T) {
	race := ParseDataPart2(inputExample)
	expected := Race{TimeAllowed: 71530, DistanceRecord: 940200}

	assert.Equal(t, expected, race)
}

func TestGetNumberOfWayToBeatRecord_Part2_1(t *testing.T) {
	race := ParseDataPart2(inputExample)
	expected := 71503

	got := race.GetNumberOfWayToBeatRecordPart2()
	assert.Equal(t, expected, got)
}

func TestGetNumberOfWayToBeatRecord_Part2(t *testing.T) {
	race := ParseDataPart2(inputExercise)
	expected := 41513103

	got := race.GetNumberOfWayToBeatRecordPart2()
	assert.Equal(t, expected, got)
}
