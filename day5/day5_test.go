package day5

import (
	"adventofcode/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestGetSeeds(t *testing.T) {
	almanac := parseData(data)

	assert.Equal(t, []int{79, 14, 55, 13}, almanac.GetSeeds())
}

func TestGetSeedsRange(t *testing.T) {
	almanac := parseData(data)
	firstSeedRange := SeedRange{From: 79, Count: 14}
	secondSeedRange := SeedRange{From: 55, Count: 13}
	seedRangesArray := []SeedRange{firstSeedRange, secondSeedRange}
	expectedSeedRanges := SeedRanges{SeedRanges: seedRangesArray}
	assert.Equal(t, expectedSeedRanges, almanac.GetSeedsRange())
}

func TestGetSoilForSeed(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 81, almanac.GetSoilForSeed(79))
	assert.Equal(t, 14, almanac.GetSoilForSeed(14))
	assert.Equal(t, 57, almanac.GetSoilForSeed(55))
	assert.Equal(t, 13, almanac.GetSoilForSeed(13))
}

func TestGetFertilizerForSoil(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 81, almanac.GetFertilizerForSoil(81))
	assert.Equal(t, 53, almanac.GetFertilizerForSoil(14))
	assert.Equal(t, 57, almanac.GetFertilizerForSoil(57))
	assert.Equal(t, 52, almanac.GetFertilizerForSoil(13))
}

func TestGetWaterForFertilizer(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 81, almanac.GetWaterForFertilizer(81))
	assert.Equal(t, 49, almanac.GetWaterForFertilizer(53))
	assert.Equal(t, 53, almanac.GetWaterForFertilizer(57))
	assert.Equal(t, 41, almanac.GetWaterForFertilizer(52))
}

func TestGetLightForWater(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 74, almanac.GetLightForWater(81))
	assert.Equal(t, 42, almanac.GetLightForWater(49))
	assert.Equal(t, 46, almanac.GetLightForWater(53))
	assert.Equal(t, 34, almanac.GetLightForWater(41))
}

func TestGetTemperatureForLight(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 78, almanac.GetTemperatureForLight(74))
	assert.Equal(t, 42, almanac.GetTemperatureForLight(42))
	assert.Equal(t, 82, almanac.GetTemperatureForLight(46))
	assert.Equal(t, 34, almanac.GetTemperatureForLight(34))
}

func TestGetHumidityForTemperature(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 78, almanac.GetHumidityForTemperature(78))
	assert.Equal(t, 43, almanac.GetHumidityForTemperature(42))
	assert.Equal(t, 82, almanac.GetHumidityForTemperature(82))
	assert.Equal(t, 35, almanac.GetHumidityForTemperature(34))
}

func TestGetLocationForHumidity(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 82, almanac.GetLocationForHumidity(78))
	assert.Equal(t, 43, almanac.GetLocationForHumidity(43))
	assert.Equal(t, 86, almanac.GetLocationForHumidity(82))
	assert.Equal(t, 35, almanac.GetLocationForHumidity(35))
}

func TestGetLowestLocation(t *testing.T) {

	almanac := parseData(data)

	assert.Equal(t, 35, almanac.GetLowestLocation())
}

func TestGetLowestLocationOfInput(t *testing.T) {

	input, err := helper.ReadFile("input")
	assert.NoError(t, err)

	almanac := parseData(input)

	assert.Equal(t, 309796150, almanac.GetLowestLocation())
}

func TestGetLowestLocationOfInputPart2(t *testing.T) {

	input, err := helper.ReadFile("input")
	assert.NoError(t, err)

	almanac := parseData(input)

	assert.Equal(t, 50716416, almanac.GetLowestLocationPart2())
}
