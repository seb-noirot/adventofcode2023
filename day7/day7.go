package day7

import (
	"sort"
	"strconv"
	"strings"
)

var strengthMap = map[string]int{
	"5":     7,
	"14":    6,
	"41":    6,
	"23":    5,
	"32":    5,
	"113":   4,
	"131":   4,
	"311":   4,
	"122":   3,
	"221":   3,
	"212":   3,
	"2111":  2,
	"1211":  2,
	"1121":  2,
	"1112":  2,
	"11111": 1,
}

var strengthCarsMap = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var strengthCarsMapPart2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 0,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

type ByHands []Hand

func (a ByHands) Len() int           { return len(a) }
func (a ByHands) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHands) Less(i, j int) bool { return CompareCards(a[i].Cards, a[j].Cards) }

type ByHands2 []Hand

func (a ByHands2) Len() int           { return len(a) }
func (a ByHands2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHands2) Less(i, j int) bool { return CompareCards2(a[i].Cards, a[j].Cards) }

func CompareCards(card1 Cards, card2 Cards) bool {
	for i := 0; i < 5; i++ {
		if strengthCarsMap[string(card1.val[i])] > strengthCarsMap[string(card2.val[i])] {
			return false
		} else if strengthCarsMap[string(card1.val[i])] < strengthCarsMap[string(card2.val[i])] {
			return true
		}
	}
	return false
}

func CompareCards2(card1 Cards, card2 Cards) bool {
	for i := 0; i < 5; i++ {
		if strengthCarsMapPart2[string(card1.val[i])] > strengthCarsMapPart2[string(card2.val[i])] {
			return false
		} else if strengthCarsMapPart2[string(card1.val[i])] < strengthCarsMapPart2[string(card2.val[i])] {
			return true
		}
	}
	return false
}

type Hands struct {
	list []Hand
}

type Hand struct {
	Cards Cards
	Bid   int
}

type Cards struct {
	val      string
	strength int
}

func ParseData(data []string) Hands {
	hands := make([]Hand, 0)
	for _, value := range data {
		hand := Hand{Cards: extractCardsFromLine(value), Bid: extractIntsFromLine(value)}
		hands = append(hands, hand)
	}
	return Hands{list: hands}
}

func ParseData2(data []string) Hands {
	hands := make([]Hand, 0)
	for _, value := range data {
		hand := Hand{Cards: extractCardsFromLine2(value), Bid: extractIntsFromLine(value)}
		hands = append(hands, hand)
	}
	return Hands{list: hands}
}

func (hands *Hands) RankBid() []int {
	handsByStrength := make(map[int][]Hand)
	bids := make([]int, 0)
	for _, hand := range hands.list {
		handsByStrength[hand.Cards.strength] = append(handsByStrength[hand.Cards.strength], hand)
	}
	for strength := 1; strength <= 7; strength++ {
		if handsByStrength[strength] == nil {
			continue
		}
		hands := RankHandsOfSameStrength(handsByStrength[strength])
		for _, hand := range hands {
			bids = append(bids, hand.Bid)
		}
	}
	return bids
}

func (hands *Hands) RankBid2() []int {
	handsByStrength := make(map[int][]Hand)
	bids := make([]int, 0)
	for _, hand := range hands.list {
		handsByStrength[hand.Cards.strength] = append(handsByStrength[hand.Cards.strength], hand)
	}
	for strength := 1; strength <= 7; strength++ {
		if handsByStrength[strength] == nil {
			continue
		}
		hands := RankHandsOfSameStrength2(handsByStrength[strength])
		for _, hand := range hands {
			bids = append(bids, hand.Bid)
		}
	}
	return bids
}

func (hands *Hands) TotalWinning() int {
	total := 0
	bids := hands.RankBid()
	for index, bid := range bids {
		total += bid * (index + 1)
	}
	return total
}

func (hands *Hands) TotalWinning2() int {
	total := 0
	bids := hands.RankBid2()
	for index, bid := range bids {
		total += bid * (index + 1)
	}
	return total
}

func RankHandsOfSameStrength(Hands []Hand) []Hand {
	sort.Sort(ByHands(Hands))
	return Hands
}

func RankHandsOfSameStrength2(Hands []Hand) []Hand {
	sort.Sort(ByHands2(Hands))
	return Hands
}

func extractCardsFromLine(line string) Cards {
	val := line[:5]
	strength := getStrengthFromVal(val)
	return Cards{val: val, strength: strength}
}

func extractCardsFromLine2(line string) Cards {
	val := line[:5]
	strength := getStrengthFromVal2(val)
	return Cards{val: val, strength: strength}
}

func getStrengthFromVal(val string) int {
	countGroupByCards := make(map[string]int)
	for _, char := range val {
		countGroupByCards[string(char)]++
	}
	counts := strings.Builder{}
	for _, count := range countGroupByCards {
		counts.WriteString(strconv.Itoa(count))
	}
	return strengthMap[counts.String()]
}

func getStrengthFromVal2(val string) int {
	countGroupByCards := make(map[string]int)
	countOfJ := 0
	for _, char := range val {
		if string(char) == "J" {
			countOfJ++
		} else {
			countGroupByCards[string(char)]++
		}
	}
	if countOfJ > 0 {
		var higherCard string
		higherCount := 0
		for card, count := range countGroupByCards {
			if higherCard == "" || count > higherCount {
				higherCard = card
				higherCount = count
			}
		}
		countGroupByCards[higherCard] += countOfJ
	}
	counts := strings.Builder{}
	for _, count := range countGroupByCards {
		counts.WriteString(strconv.Itoa(count))
	}
	return strengthMap[counts.String()]
}

func extractIntsFromLine(line string) int {
	bidPart := line[6:]
	val, _ := strconv.Atoi(bidPart)
	return val
}
