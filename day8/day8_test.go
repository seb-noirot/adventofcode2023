package day8

import (
	"adventofcode/helper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var exampleData = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

var exampleData2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

var exampleData3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestParseData(t *testing.T) {
	data := ParseData(exampleData)
	assert.Equal(t, "AAA", data.Value)
	assert.Equal(t, "BBB", data.Left.Value)
	assert.Equal(t, "CCC", data.Right.Value)
}

func TestParseData2(t *testing.T) {
	data := ParseData(exampleData2)
	assert.Equal(t, "AAA", data.Value)
	assert.Equal(t, "BBB", data.Left.Value)
	assert.Equal(t, "BBB", data.Right.Value)
}

func TestTraverseTree(t *testing.T) {
	data := ParseData(exampleData)
	assert.Equal(t, 2, traverseTree(data, "RL"))
}

func TestTraverseTree2(t *testing.T) {
	data := ParseData(exampleData2)
	assert.Equal(t, 6, traverseTree(data, "LLR"))
}

func TestTraverseTree3(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)

	data := ParseData(content)
	instructions := strings.Split(content, "\n")[0]
	assert.Equal(t, 242, traverseTree(data, instructions))
}

func TestParseDataWithEndA(t *testing.T) {
	data := ParseDataWithEndA(exampleData3)
	assert.Equal(t, 2, len(data))
}

func TestTraverseTreeWithEndA(t *testing.T) {
	data := ParseDataWithEndA(exampleData3)
	assert.Equal(t, 6, traverseTreeWithMultipleNode(data, "LR"))
}

func TestTraverseTree_Part2(t *testing.T) {
	content, err := helper.ReadFile("input")
	assert.NoError(t, err)

	data := ParseDataWithEndA(content)
	instructions := strings.Split(content, "\n")[0]
	assert.Equal(t, 8245452805243, traverseTreeWithMultipleNode(data, instructions))
}
