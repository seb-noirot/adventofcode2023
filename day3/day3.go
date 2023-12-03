package day3

import (
	"strconv"
	"strings"
	"unicode"
)

func sumOfPartNumber(inputs []string) int {
	sum := 0
	matrix := toMatrix(inputs)
	for rowIndex, row := range matrix {
		adjacentToSymbol := false
		currentNumber := strings.Builder{}
		for cellIndex, cell := range row {
			if unicode.IsDigit(cell) {
				currentNumber.WriteRune(cell)
				adj := isAdjacentToSymbol(rowIndex, cellIndex, matrix)

				if !adjacentToSymbol && adj {
					adjacentToSymbol = true
				}
			} else {
				if adjacentToSymbol {
					value, _ := strconv.Atoi(currentNumber.String())
					sum += value
				}
				currentNumber.Reset()
				adjacentToSymbol = false
			}
		}
		if adjacentToSymbol {
			value, _ := strconv.Atoi(currentNumber.String())
			sum += value
		}
	}
	return sum
}

func sumOfPartGearsRatio(inputs []string) int {
	matrix := toMatrix(inputs)
	// Find all numbers
	allRatios := findAllNumbers(matrix)
	// Find all gears
	allGears := findAllGears(matrix)
	// Find all ratios
	gearsWithRatios := findGearsWithRatios(allGears, allRatios)
	// Multiply all ratios
	sum := sumAllGearsWith2Rations(gearsWithRatios)
	return sum
}

type Ratio struct {
	RatioId  int
	RowIndex int
	Cells    []Cell
	Value    int
}

type Cell struct {
	RowIndex  int
	CellIndex int
}

type Gear struct {
	Cell Cell
}

type RatiosMap struct {
	Ratios map[int]Ratio
}

func sumAllGearsWith2Rations(ratios map[Gear]RatiosMap) int {
	sum := 0
	for _, ratiosMap := range ratios {
		if len(ratiosMap.Ratios) == 2 {
			sum += multiplyRatios(ratiosMap)
		}
	}
	return sum
}

func multiplyRatios(ratios RatiosMap) int {
	product := 1
	for _, ratio := range ratios.Ratios {
		product *= ratio.Value
	}
	return product
}

func findGearsWithRatios(gears []Gear, ratios []Ratio) map[Gear]RatiosMap {
	gearsWithRatios := make(map[Gear]RatiosMap)
	for _, gear := range gears {
		ratiosMap := RatiosMap{Ratios: make(map[int]Ratio)}
		for _, ratio := range ratios {
			for _, cell := range ratio.Cells {
				if isCloseCell(gear.Cell, cell) {
					ratiosMap.Ratios[ratio.RatioId] = ratio
				}
			}
		}
		gearsWithRatios[gear] = ratiosMap
	}
	return gearsWithRatios
}

func findAllGears(matrix [][]rune) []Gear {
	var gears []Gear
	for rowIndex, row := range matrix {
		for cellIndex, cell := range row {
			if cell == '*' {
				gears = append(gears, Gear{Cell{rowIndex, cellIndex}})
			}
		}
	}
	return gears
}

func isCloseCell(cell Cell, otherCell Cell) bool {
	if cell.RowIndex == otherCell.RowIndex {
		if cell.CellIndex == otherCell.CellIndex+1 || cell.CellIndex == otherCell.CellIndex-1 {
			return true
		}
	} else if cell.RowIndex == otherCell.RowIndex-1 {
		if cell.CellIndex == otherCell.CellIndex || cell.CellIndex == otherCell.CellIndex+1 || cell.CellIndex == otherCell.CellIndex-1 {
			return true
		}
	} else if cell.RowIndex == otherCell.RowIndex+1 {
		if cell.CellIndex == otherCell.CellIndex || cell.CellIndex == otherCell.CellIndex+1 || cell.CellIndex == otherCell.CellIndex-1 {
			return true
		}
	}
	return false
}

func findAllNumbers(matrix [][]rune) []Ratio {
	var ratios []Ratio
	ratioId := 0
	for rowIndex, row := range matrix {
		found := false
		var ratio Ratio
		val := strings.Builder{}
		for cellIndex, cell := range row {
			if unicode.IsDigit(cell) {
				if !found {
					found = true
					ratio = Ratio{}
					ratio.Cells = make([]Cell, 0)
					ratio.RatioId = ratioId
					ratioId++
				}
				ratio.Cells = append(ratio.Cells, Cell{rowIndex, cellIndex})
				val.WriteRune(cell)
			} else {
				if found {
					value, _ := strconv.Atoi(val.String())
					ratio.Value = value
					ratios = append(ratios, ratio)
					val.Reset()
					found = false
				}
			}
		}
		if found {
			value, _ := strconv.Atoi(val.String())
			ratio.Value = value
			ratios = append(ratios, ratio)
		}
	}
	return ratios
}

func toMatrix(inputs []string) [][]rune {
	matrix := make([][]rune, len(inputs))
	for rowIndex, row := range inputs {
		matrix[rowIndex] = make([]rune, len(row))
		for cellIndex, cell := range row {
			matrix[rowIndex][cellIndex] = cell
		}
	}
	return matrix
}

func isAdjacentToSymbol(rowIndex int, cellIndex int, matrix [][]rune) bool {
	// look at all 8 directions
	// top

	if rowIndex > 0 {
		if isSymbol(matrix[rowIndex-1][cellIndex]) {
			return true
		}
	}
	// bottom
	if rowIndex < len(matrix)-1 {
		if isSymbol(matrix[rowIndex+1][cellIndex]) {
			return true
		}
	}
	// left
	if cellIndex > 0 {
		if isSymbol(matrix[rowIndex][cellIndex-1]) {
			return true
		}

	}
	// right
	if cellIndex < len(matrix[rowIndex])-1 {
		if isSymbol(matrix[rowIndex][cellIndex+1]) {
			return true
		}
	}
	// top-left
	if rowIndex > 0 && cellIndex > 0 {
		if isSymbol(matrix[rowIndex-1][cellIndex-1]) {
			return true
		}
	}
	// top-right
	if rowIndex > 0 && cellIndex < len(matrix[rowIndex])-1 {
		if isSymbol(matrix[rowIndex-1][cellIndex+1]) {
			return true
		}

	}

	// bottom-left
	if rowIndex < len(matrix)-1 && cellIndex > 0 {
		if isSymbol(matrix[rowIndex+1][cellIndex-1]) {
			return true
		}

	}

	// bottom-right
	if rowIndex < len(matrix)-1 && cellIndex < len(matrix[rowIndex])-1 {
		if isSymbol(matrix[rowIndex+1][cellIndex+1]) {
			return true
		}
	}
	return false
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}
