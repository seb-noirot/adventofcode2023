package day7

import (
	"adventofcode/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483"}

func TestParseData(t *testing.T) {
	got := ParseData(testInput)
	expected := Hands{list: []Hand{
		{Cards: Cards{val: "32T3K", strength: 2}, Bid: 765},
		{Cards: Cards{val: "T55J5", strength: 4}, Bid: 684},
		{Cards: Cards{val: "KK677", strength: 3}, Bid: 28},
		{Cards: Cards{val: "KTJJT", strength: 3}, Bid: 220},
		{Cards: Cards{val: "QQQJA", strength: 4}, Bid: 483}}}

	assert.Equal(t, expected, got)
}

func TestParseData2(t *testing.T) {
	got := ParseData2(testInput)
	expected := Hands{list: []Hand{
		{Cards: Cards{val: "32T3K", strength: 2}, Bid: 765},
		{Cards: Cards{val: "T55J5", strength: 6}, Bid: 684},
		{Cards: Cards{val: "KK677", strength: 3}, Bid: 28},
		{Cards: Cards{val: "KTJJT", strength: 6}, Bid: 220},
		{Cards: Cards{val: "QQQJA", strength: 6}, Bid: 483}}}

	assert.Equal(t, expected, got)
}

func TestRankBid(t *testing.T) {
	hands := ParseData(testInput)
	expected := []int{765, 220, 28, 684, 483}

	got := hands.RankBid()
	assert.Equal(t, expected, got)
}

func TestTotalWinning(t *testing.T) {
	hands := ParseData(testInput)
	expected := 6440

	got := hands.TotalWinning()
	assert.Equal(t, expected, got)
}

func TestTotalWinningPart2Demo(t *testing.T) {
	hands := ParseData2(testInput)
	expected := 5905

	got := hands.TotalWinning2()
	assert.Equal(t, expected, got)
}

func TestTotalWinningPart1(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	hands := ParseData(input)
	expected := 248453531

	got := hands.TotalWinning()
	assert.Equal(t, expected, got)
}

func TestTotalWinningPart2(t *testing.T) {
	input, err := helper.ReadLines("input")
	assert.NoError(t, err)

	hands := ParseData2(input)
	expected := 248781813

	got := hands.TotalWinning2()
	assert.Equal(t, expected, got)
}
