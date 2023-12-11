package day11

import (
	"strconv"
	"strings"
)

type Universe struct {
	Dimension Dimension
	Galaxys   []Galaxy
	Age       int
}

type Dimension struct {
	X int
	Y int
}

type Galaxy struct {
	Name        string
	Coordinates Coordinates
}

type Coordinates struct {
	X int
	Y int
}

func ParseData(data string, ageUniverse int) Universe {
	lines := strings.Split(data, "\n")
	galaxys := make([]Galaxy, 0)
	count := 0
	for y, line := range lines {
		for x, val := range line {
			if val == '.' {
				continue
			}
			galaxy := Galaxy{
				Name:        strconv.Itoa(count + 1),
				Coordinates: Coordinates{X: x, Y: y},
			}
			galaxys = append(galaxys, galaxy)
			count++
		}
	}
	yDimension := len(lines)
	xDimension := len(lines[0])
	dimension := Dimension{X: xDimension, Y: yDimension}
	return Universe{Dimension: dimension, Galaxys: galaxys, Age: ageUniverse}
}

func (universe Universe) ExpendDimension() Universe {
	dimension := universe.Dimension
	rowWithoutGalaxy, colWithoutGalaxy := universe.getPartOfTheUniverseWithoutGalaxy()
	newGalaxys := universe.getNewGalaxys(rowWithoutGalaxy, colWithoutGalaxy)
	newDimension := Dimension{X: dimension.X + len(colWithoutGalaxy)*universe.Age, Y: dimension.Y + len(rowWithoutGalaxy)*universe.Age}
	return Universe{Dimension: newDimension, Galaxys: newGalaxys}
}

func (universe Universe) getNewGalaxys(rowWithoutGalaxy []int, colWithoutGalaxy []int) []Galaxy {
	mapGalaxies := make(map[string]Galaxy)
	for _, galaxy := range universe.Galaxys {
		mapGalaxies[galaxy.Name] = galaxy
	}
	for _, galaxy := range universe.Galaxys {
		expensionX := 0
		expensionY := 0
		for _, row := range rowWithoutGalaxy {
			if galaxy.Coordinates.Y > row {
				expensionY += universe.Age - 1
			}
		}
		for _, col := range colWithoutGalaxy {
			if galaxy.Coordinates.X > col {
				expensionX += universe.Age - 1
			}
		}
		galaxy.Coordinates.X += expensionX
		galaxy.Coordinates.Y += expensionY
		mapGalaxies[galaxy.Name] = galaxy
	}
	newGalaxys := make([]Galaxy, len(universe.Galaxys))
	for name, galaxy := range mapGalaxies {
		index, _ := strconv.Atoi(name)
		newGalaxys[index-1] = galaxy
	}
	return newGalaxys
}

func (universe Universe) getPartOfTheUniverseWithoutGalaxy() ([]int, []int) {
	dimension := universe.Dimension
	rowWithoutGalaxy := make([]int, 0)
	colWithoutGalaxy := make([]int, 0)
	for row := 0; row < dimension.Y; row++ {
		if !universe.anyGalaxyInRow(row) {
			rowWithoutGalaxy = append(rowWithoutGalaxy, row)
		}
	}
	for col := 0; col < dimension.X; col++ {
		if !universe.anyGalaxyInCol(col) {
			colWithoutGalaxy = append(colWithoutGalaxy, col)
		}
	}
	return rowWithoutGalaxy, colWithoutGalaxy
}

func (universe Universe) anyGalaxyInRow(row int) bool {
	for _, galaxy := range universe.Galaxys {
		if galaxy.Coordinates.Y == row {
			return true
		}
	}
	return false
}

func (universe Universe) anyGalaxyInCol(col int) bool {
	for _, galaxy := range universe.Galaxys {
		if galaxy.Coordinates.X == col {
			return true
		}
	}
	return false
}

func (universe Universe) CalculateSumShortestPathBetweenGalaxies() int {
	sum := 0
	for _, galaxy := range universe.Galaxys {
		sum += calculateSumShortestPathWithOtherGalaxies(galaxy, universe.Galaxys)
	}
	return sum / 2 // because we calculate the distance between 2 galaxies twice
}

func calculateSumShortestPathWithOtherGalaxies(galaxy Galaxy, galaxys []Galaxy) int {
	sum := 0
	for _, otherGalaxy := range galaxys {
		if galaxy.Name == otherGalaxy.Name {
			continue
		}
		distance := calculateShortestPath(galaxy, otherGalaxy)
		sum += distance
	}
	return sum
}

func calculateShortestPath(galaxy Galaxy, galaxy2 Galaxy) int {
	return absInt(galaxy.Coordinates.X-galaxy2.Coordinates.X) + absInt(galaxy.Coordinates.Y-galaxy2.Coordinates.Y)
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
