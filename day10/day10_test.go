package day10

import (
	"adventofcode/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleData = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func TestParseData(t *testing.T) {
	data := ParseData(exampleData)
	assert.Equal(t, 5, len(data.Rows))
	assert.Equal(t, 5, len(data.Rows[0]))
	assert.Equal(t, 5, len(data.Rows[1]))
	assert.Equal(t, 5, len(data.Rows[2]))
	assert.Equal(t, 5, len(data.Rows[3]))
	assert.Equal(t, 5, len(data.Rows[4]))

	fmt.Printf("%+v\n", data)
}

func TestGetStartingPoint(t *testing.T) {
	data := ParseData(exampleData)
	x, y := data.GetStartingPoint()
	assert.Equal(t, 0, x)
	assert.Equal(t, 2, y)
}

func TestCalculatePath(t *testing.T) {
	data := ParseData(exampleData)
	furtherPoint := data.FollowPipe("F")
	assert.Equal(t, 8, len(furtherPoint)/2)
}

func TestCalculatePathPart1(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)
	data := ParseData(content)
	furtherPoint := data.FollowPipe("7")
	assert.Equal(t, 6956, len(furtherPoint)/2)
}

func TestCalculateInLoop(t *testing.T) {
	content, err := helper.ReadFile("input2Demo")
	assert.NoError(t, err)
	data := ParseData(content)
	enclosedInLoop := data.EnclosedInLoop("F")
	assert.Equal(t, 8, enclosedInLoop)
}

func TestCalculateInLoopPart2(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)
	data := ParseData(content)
	enclosedInLoop := data.EnclosedInLoop("7")
	assert.Equal(t, 8, enclosedInLoop)
}

func TestCalculateInLoop_InputDemoSimple(t *testing.T) {
	content, err := helper.ReadFile("inputDemoSimple")
	assert.NoError(t, err)
	data := ParseData(content)
	enclosedInLoop := data.EnclosedInLoop("F")
	assert.Equal(t, 4, enclosedInLoop)
}

func TestCalculateInLoop_InputDemoSimple2(t *testing.T) {
	content, err := helper.ReadFile("inputDemoSimple2")
	assert.NoError(t, err)
	data := ParseData(content)
	enclosedInLoop := data.EnclosedInLoop("F")
	assert.Equal(t, 4, enclosedInLoop)
}
