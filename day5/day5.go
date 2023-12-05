package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds                 []int
	SeedsRange            SeedRanges
	SeedToSoil            Intervals
	SoilToFertilizer      Intervals
	FertilizerToWater     Intervals
	WaterToLight          Intervals
	LightToTemperature    Intervals
	TemperatureToHumidity Intervals
	HumidityToLocation    Intervals
}

type SeedRanges struct {
	SeedRanges []SeedRange
}

type SeedRange struct {
	From  int
	Count int
}

type Intervals struct {
	Intervals []Interval
}

type Interval struct {
	From  int
	To    int
	Count int
}

func (intervals *Intervals) GetFrom(fromInput int) (int, bool) {
	for _, interval := range intervals.Intervals {
		to, ok := interval.GetFrom(fromInput)
		if ok {
			return to, ok
		}
	}
	return fromInput, false
}

func (interval *Interval) GetFrom(fromInput int) (int, bool) {
	if fromInput >= interval.From && fromInput < interval.From+interval.Count {
		delta := fromInput - interval.From
		return interval.To + delta, true
	}
	return fromInput, false
}

func (Almanac *Almanac) GetSeedsRange() SeedRanges {
	return Almanac.SeedsRange
}

func (Almanac *Almanac) GetSeeds() []int {
	return Almanac.Seeds
}

func (almanac *Almanac) GetSoilForSeed(seed int) int {
	soil, ok := almanac.SeedToSoil.GetFrom(seed)
	if !ok {
		return seed
	}
	return soil
}

func (almanac *Almanac) GetFertilizerForSoil(soil int) int {
	fertilizer, ok := almanac.SoilToFertilizer.GetFrom(soil)
	if !ok {
		return soil
	}
	return fertilizer
}

func (almanac *Almanac) GetWaterForFertilizer(fertilizer int) int {
	water, ok := almanac.FertilizerToWater.GetFrom(fertilizer)
	if !ok {
		return fertilizer
	}
	return water
}

func (almanac *Almanac) GetLightForWater(water int) int {
	light, ok := almanac.WaterToLight.GetFrom(water)
	if !ok {
		return water
	}
	return light
}

func (almanac *Almanac) GetTemperatureForLight(light int) int {
	temperature, ok := almanac.LightToTemperature.GetFrom(light)
	if !ok {
		return light
	}
	return temperature
}

func (almanac *Almanac) GetHumidityForTemperature(temperature int) int {
	humidity, ok := almanac.TemperatureToHumidity.GetFrom(temperature)
	if !ok {
		return temperature
	}
	return humidity
}

func (almanac *Almanac) GetLocationForHumidity(humidity int) int {
	location, ok := almanac.HumidityToLocation.GetFrom(humidity)
	if !ok {
		return humidity
	}
	return location
}

func (almanc *Almanac) GetHighestLocation() int {
	maxLocation := -1
	for _, seed := range almanc.Seeds {
		seedLocation := almanc.GetLocationForSeed(seed)
		if seedLocation > maxLocation {
			maxLocation = seedLocation
		}
	}
	return maxLocation
}

func (almanac *Almanac) GetLowestLocation() int {
	minLocation := -1
	for _, seed := range almanac.Seeds {
		seedLocation := almanac.GetLocationForSeed(seed)
		if minLocation == -1 || seedLocation < minLocation {
			minLocation = seedLocation
		}
	}
	return minLocation
}

func (almanac *Almanac) GetLowestLocationPart2() int {
	minLocation := -1
	for index, seedRange := range almanac.SeedsRange.SeedRanges {
		fmt.Printf("index: %d, seedRange: %+v\n", index, seedRange)
		seedLocation := almanac.GetLowestLocationSeedRange(seedRange)
		if minLocation == -1 || seedLocation < minLocation {
			minLocation = seedLocation
		}
	}
	return minLocation
}

func (almanac *Almanac) GetLowestLocationSeedRange(seedRange SeedRange) int {
	minLocation := -1
	fmt.Printf("total seeds: %d\n", seedRange.Count)
	count := 0
	for seed := seedRange.From; seed < seedRange.From+seedRange.Count; seed++ {
		count++
		if (count % 10_000_000) == 0 {
			fmt.Printf("count: %d\n", count)
		}
		seedLocation := almanac.GetLocationForSeed(seed)
		if minLocation == -1 || seedLocation < minLocation {
			minLocation = seedLocation
		}
	}
	return minLocation
}

func (almanac *Almanac) GetLocationForSeed(seed int) int {
	soil := almanac.GetSoilForSeed(seed)
	fertilizer := almanac.GetFertilizerForSoil(soil)
	water := almanac.GetWaterForFertilizer(fertilizer)
	light := almanac.GetLightForWater(water)
	temperature := almanac.GetTemperatureForLight(light)
	humidity := almanac.GetHumidityForTemperature(temperature)
	location := almanac.GetLocationForHumidity(humidity)
	return location
}

func parseData(data string) Almanac {

	lines := strings.Split(data, "\n")
	seeds := parseSeeds(lines)
	seedRanges := parseSeedsRange(lines)
	seedToSoil := extractMap(lines, "seed-to-soil map:")
	soilToFertilizer := extractMap(lines, "soil-to-fertilizer map:")
	fertilizerToWater := extractMap(lines, "fertilizer-to-water map:")
	waterToLight := extractMap(lines, "water-to-light map:")
	lightToTemperature := extractMap(lines, "light-to-temperature map:")
	temperatureToHumidity := extractMap(lines, "temperature-to-humidity map:")
	humidityToLocation := extractMap(lines, "humidity-to-location map:")
	almanac := Almanac{
		Seeds:                 seeds,
		SeedsRange:            seedRanges,
		SeedToSoil:            seedToSoil,
		SoilToFertilizer:      soilToFertilizer,
		FertilizerToWater:     fertilizerToWater,
		WaterToLight:          waterToLight,
		LightToTemperature:    lightToTemperature,
		TemperatureToHumidity: temperatureToHumidity,
		HumidityToLocation:    humidityToLocation,
	}
	return almanac
}

func extractMap(lines []string, header string) Intervals {
	foundHeader := false
	intervals := Intervals{Intervals: make([]Interval, 0)}
	for _, line := range lines {
		if line == header {
			foundHeader = true
			continue
		} else if foundHeader {
			if strings.Trim(line, " ") == "" {
				return intervals
			} else {
				interval := parseLine(line)
				intervals = appendInterval(intervals, interval)
			}
		}
	}
	return intervals

}

func appendInterval(intervals Intervals, interval Interval) Intervals {
	newIntervals := append(intervals.Intervals, interval)
	return Intervals{Intervals: newIntervals}
}

func parseLine(line string) Interval {
	vals := strings.Split(line, " ")
	to, _ := strconv.Atoi(vals[0])
	from, _ := strconv.Atoi(vals[1])
	count, _ := strconv.Atoi(vals[2])

	return Interval{From: from, To: to, Count: count}
}

func parseSeeds(lines []string) []int {
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds: ") {
			values := line[7:]
			return parseSeedLine(values)
		}
	}
	return []int{}
}

func parseSeedsRange(lines []string) SeedRanges {
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds: ") {
			values := line[7:]
			return parseSeedLineRanges(values)
		}
	}
	return SeedRanges{}
}

func parseSeedLineRanges(line string) SeedRanges {
	seedRanges := SeedRanges{}
	seedRange := SeedRange{}
	for index, seedRangeVal := range strings.Split(line, " ") {
		if index%2 == 0 {
			from, _ := strconv.Atoi(seedRangeVal)
			seedRange.From = from
		} else {
			count, _ := strconv.Atoi(seedRangeVal)
			seedRange.Count = count
			seedRanges.SeedRanges = append(seedRanges.SeedRanges, seedRange)
		}
	}
	return seedRanges
}

func parseSeedLine(line string) []int {
	seeds := []int{}
	for _, seedVal := range strings.Split(line, " ") {
		seed, _ := strconv.Atoi(seedVal)
		seeds = append(seeds, seed)
	}
	return seeds
}
