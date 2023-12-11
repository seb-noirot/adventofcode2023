package day10

import (
	"fmt"
	"strings"
)

var Left = Direction{X: -1, Y: 0}
var Right = Direction{X: 1, Y: 0}
var Up = Direction{X: 0, Y: -1}
var Down = Direction{X: 0, Y: 1}
var Stay = Direction{X: 0, Y: 0}
var Unknown = Direction{X: -2, Y: -2}

var instructionsMap = map[string]Instructions{
	"|": {Name: "|", From: Up, To: Down},
	"-": {Name: "-", From: Left, To: Right},
	"L": {Name: "L", From: Up, To: Right},
	"J": {Name: "J", From: Left, To: Up},
	"7": {Name: "7", From: Left, To: Down},
	"F": {Name: "F", From: Right, To: Down},
	".": {Name: ".", From: Stay, To: Stay},
	"S": {Name: "S", From: Stay, To: Stay},
}

type Cell struct {
	X int
	Y int
}
type InstructionStep struct {
	Instructions Instructions
	Steps        int
}

type Instructions struct {
	Name string
	From Direction
	To   Direction
}

type Direction struct {
	X int
	Y int
}

type Matrix struct {
	Rows [][]string
}

type Position struct {
	previousX int
	previousY int
	currentX  int
	currentY  int
	nextX     int
	nextY     int
}

func ParseData(data string) Matrix {
	matrix := Matrix{Rows: make([][]string, 0)}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		matrix.Rows = append(matrix.Rows, ParseDataLine(line))
	}
	return matrix
}

func ParseDataLine(line string) []string {
	row := make([]string, 0)
	for _, val := range strings.Split(line, "") {
		if strings.Trim(val, " ") == "" {
			continue
		}
		row = append(row, val)
	}
	return row
}

func (matrix Matrix) GetStartingPoint() (int, int) {
	for y, row := range matrix.Rows {
		for x, val := range row {
			if val == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}

func (matrix Matrix) FollowPipe(key string) []Position {
	x, y := matrix.GetStartingPoint()
	instruction := instructionsMap[key]
	fmt.Printf("Starting at: %d, %d\n", x, y)
	fmt.Printf("Instruction: %+v\n", instruction)
	fmt.Printf("Instruction: %+v\n", instruction.From)
	position := Position{previousX: x, previousY: y, currentX: x, currentY: y, nextX: x + instruction.From.X, nextY: y + instruction.From.Y}
	path := make([]Position, 0)
	count := 0
	for {
		position = matrix.FollowPipeStep(position)
		path = append(path, position)
		if position.isBackToStartingPoint(x, y) {
			break
		}
		count++
	}
	return path
}

func (position Position) isBackTo() bool {
	return position.previousX == position.nextX && position.previousY == position.nextY
}

func (position Position) isBackToStartingPoint(x int, y int) bool {
	return position.currentX == x && position.currentY == y
}

func (matrix Matrix) FollowPipeStep(position Position) Position {
	nextCell := matrix.Rows[position.nextY][position.nextX]
	instructions := instructionsMap[nextCell]

	instructionFrom := instructions.From
	positionFrom := Position{position.currentX, position.currentY, position.nextX, position.nextY, position.nextX + instructionFrom.X, position.nextY + instructionFrom.Y}
	if !positionFrom.isBackTo() {
		return positionFrom
	} else {
		instructionTo := instructions.To
		positionTo := Position{position.currentX, position.currentY, position.nextX, position.nextY, position.nextX + instructionTo.X, position.nextY + instructionTo.Y}
		return positionTo
	}
}

func (matrix Matrix) EnclosedInLoop(key string) int {
	loopPath := matrix.FollowPipe(key)
	return matrix.FindCellInTheLoop(loopPath)
}

func (matrix Matrix) FindCellInTheLoop(path []Position) int {
	loopPath := make([]Cell, 0)
	for _, position := range path {
		cell := Cell{X: position.currentX, Y: position.currentY}
		loopPath = append(loopPath, cell)
	}
	newMatrix := matrix.CreateMatrixWithLoop(loopPath)
	printMatrix(newMatrix)
	markCellNotInLoop(newMatrix)
	fmt.Println("")
	fmt.Println("")

	printMatrix(newMatrix)

	return countTheDot(newMatrix)
}

func countTheDot(matrix [][]string) int {
	count := 0
	for _, row := range matrix {
		for _, val := range row {
			if val == "." {
				count++
			}
		}
	}
	return count

}

func markCellNotInLoop(matrix [][]string) {
	// From each cell on the border

	marked := true
	for marked {
		marked = false
		for i := 0; i < len(matrix); i++ {
			row := matrix[i]
			newMarked := markRow(i, row, matrix)
			if newMarked {
				marked = true
			}
		}

		for i := len(matrix) - 1; i >= 0; i-- {
			row := matrix[i]
			newMarked := markRow(i, row, matrix)
			if newMarked {
				marked = true
			}
		}
	}
}

func markRow(y int, row []string, matrix [][]string) bool {
	marked := false
	if y == 0 {
		for i := 0; i < len(row); i++ {
			if row[i] == "." {
				markCell(i, y, matrix)
				marked = true
			}
		}
	} else {
		for i := 0; i < len(row); i++ {
			cell := row[i]
			if cell != "." {
				continue
			}
			if cell == "0" {
				continue
			}
			if isInContactOfX(i, y, matrix) || isExclude(i, y, matrix) {
				markCell(i, y, matrix)
				marked = true
			}
		}
	}
	return marked
}

func isExclude(x int, y int, matrix [][]string) bool {
	return false
}

func isInContactOfX(x int, y int, matrix [][]string) bool {
	if y == 0 || y == len(matrix)-1 || x == 0 || x == len(matrix[y])-1 {
		return true
	}
	// Check up
	cellUp := matrix[y-1][x]
	if cellUp == "0" {
		return true
	}
	// Check down
	cellDown := matrix[y+1][x]
	if cellDown == "0" {
		return true
	}
	// Check left
	cellLeft := matrix[y][x-1]
	if cellLeft == "0" {
		return true
	}
	// Check right
	cellRight := matrix[y][x+1]
	if cellRight == "0" {
		return true
	}
	return false
}

func markCell(x int, y int, matrix [][]string) {
	matrix[y][x] = "0"
}
func (matrix Matrix) CreateMatrixWithLoop(loopPath []Cell) [][]string {
	newMatrix := make([][]string, 0)
	for _, row := range matrix.Rows {
		newRow := make([]string, len(row))
		for i := range newRow {
			newRow[i] = "."
		}
		newMatrix = append(newMatrix, newRow)
	}
	for _, cell := range loopPath {
		newMatrix[cell.Y][cell.X] = matrix.Rows[cell.Y][cell.X]
	}
	return newMatrix
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}

func printMatrixRune(matrix [][]rune) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%v ", val)
		}
		fmt.Println()
	}
}
