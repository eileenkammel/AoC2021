package aoc

import (
	"log"
	"strconv"
	"strings"
)

type BingoInput []int

type BingoCard struct {
	values    [5][5]int
	filled_in [5][5]bool
}

func NewBingoCard(values [5][5]int) BingoCard {
	return BingoCard{
		values: values,
	}
}

func (c *BingoCard) FillInValue(value int) {
	for x := range c.values {
		for y := range c.values[x] {
			if c.values[x][y] == value {
				c.filled_in[x][y] = true
			}
		}
	}
}

func (c *BingoCard) NotFilledSum() int {
	var sum int
	for x := range c.values {
		for y := range c.values[x] {
			if !c.filled_in[x][y] {
				sum += c.values[x][y]
			}
		}
	}
	return sum
}

func (c *BingoCard) HasBingo() bool {
	for x := range c.filled_in {
		if c.rowHasBingo(x) || c.colHasBingo(x) {
			return true
		}
	}
	return false
}

func (c *BingoCard) rowHasBingo(rowNum int) bool {
	for x := range c.filled_in {
		if !c.filled_in[x][rowNum] {
			return false
		}
	}
	return true
}

func (c *BingoCard) colHasBingo(colNum int) bool {
	for x := range c.filled_in {
		if !c.filled_in[colNum][x] {
			return false
		}
	}
	return true
}

// parseInput expects s to be a string of comma seperated integers
func parseInput(s string) BingoInput {
	input := BingoInput{}
	for _, numberString := range strings.Split(s, ",") {
		number, _ := strconv.Atoi(numberString)
		input = append(input, number)
	}
	return input
}

// parseBingoCard expects 5 strings where each contains 5 number
// seprated by space
func parseBingoCard(s []string) BingoCard {
	card := BingoCard{}
	for row, currentLine := range s {
		values := strings.Fields(currentLine)
		for column, numberString := range values {
			number, _ := strconv.Atoi(numberString)
			card.values[row][column] = number
		}
	}
	return card
}

func ParseChallange(s string) (BingoInput, []BingoCard) {
	inputLines := strings.Split(s, "\n")
	bingoInput := parseInput(inputLines[0])

	inputLines = inputLines[1:]
	if len(inputLines)%6 != 0 {
		log.Fatalf("Number of remaining lines broken")
	}

	var cards []BingoCard
	for len(inputLines) != 0 {
		inputLines = inputLines[1:]
		cards = append(cards, parseBingoCard(inputLines[:5]))
		inputLines = inputLines[5:]
	}

	return bingoInput, cards
}
