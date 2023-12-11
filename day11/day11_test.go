package day11

import (
	"adventofcode/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var inputDemo = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestParseData(t *testing.T) {
	universe := ParseData(inputDemo, 1)
	dimension := universe.Dimension
	expectedDimension := Dimension{X: 10, Y: 10}
	assert.Equal(t, expectedDimension, dimension)
	galaxys := universe.Galaxys
	expectedGalaxys := []Galaxy{
		{Name: "1", Coordinates: Coordinates{X: 3, Y: 0}},
		{Name: "2", Coordinates: Coordinates{X: 7, Y: 1}},
		{Name: "3", Coordinates: Coordinates{X: 0, Y: 2}},
		{Name: "4", Coordinates: Coordinates{X: 6, Y: 4}},
		{Name: "5", Coordinates: Coordinates{X: 1, Y: 5}},
		{Name: "6", Coordinates: Coordinates{X: 9, Y: 6}},
		{Name: "7", Coordinates: Coordinates{X: 7, Y: 8}},
		{Name: "8", Coordinates: Coordinates{X: 0, Y: 9}},
		{Name: "9", Coordinates: Coordinates{X: 4, Y: 9}},
	}
	assert.Equal(t, expectedGalaxys, galaxys)
}

func TestExpendDimension(t *testing.T) {
	universe := ParseData(inputDemo, 1)
	universe = universe.ExpendDimension()
	dimension := universe.Dimension
	expectedDimension := Dimension{X: 13, Y: 12}
	assert.Equal(t, expectedDimension, dimension)
	galaxys := universe.Galaxys
	expectedGalaxys := []Galaxy{
		{Name: "1", Coordinates: Coordinates{X: 4, Y: 0}},
		{Name: "2", Coordinates: Coordinates{X: 9, Y: 1}},
		{Name: "3", Coordinates: Coordinates{X: 0, Y: 2}},
		{Name: "4", Coordinates: Coordinates{X: 8, Y: 5}},
		{Name: "5", Coordinates: Coordinates{X: 1, Y: 6}},
		{Name: "6", Coordinates: Coordinates{X: 12, Y: 7}},
		{Name: "7", Coordinates: Coordinates{X: 9, Y: 10}},
		{Name: "8", Coordinates: Coordinates{X: 0, Y: 11}},
		{Name: "9", Coordinates: Coordinates{X: 5, Y: 11}},
	}
	assert.Equal(t, expectedGalaxys, galaxys)
}

func TestCalculateSumShortestPathBetweenGalaxies(t *testing.T) {
	universe := ParseData(inputDemo, 1)
	universe = universe.ExpendDimension()
	sumShortestPath := universe.CalculateSumShortestPathBetweenGalaxies()
	assert.Equal(t, 374, sumShortestPath)
}

func TestCalculateSumShortestPathBetweenGalaxies_10Year(t *testing.T) {
	universe := ParseData(inputDemo, 10)
	universe = universe.ExpendDimension()
	sumShortestPath := universe.CalculateSumShortestPathBetweenGalaxies()
	assert.Equal(t, 1030, sumShortestPath)
}

func TestCalculateSumShortestPathBetweenGalaxies_100Year(t *testing.T) {
	universe := ParseData(inputDemo, 100)
	universe = universe.ExpendDimension()
	sumShortestPath := universe.CalculateSumShortestPathBetweenGalaxies()
	assert.Equal(t, 8410, sumShortestPath)
}

func TestCalculateSumShortestPathBetweenGalaxiesPart2(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)
	universe := ParseData(content, 1000000)
	universe = universe.ExpendDimension()
	sumShortestPath := universe.CalculateSumShortestPathBetweenGalaxies()
	assert.Equal(t, 742305960572, sumShortestPath)
}
