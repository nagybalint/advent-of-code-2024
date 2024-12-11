package tasks

import (
	"strconv"
	"strings"
)

type Day11Task1 struct{}
type Day11Task2 struct{}
type Day11 struct{}
type stone string

var stoneChangeHelper = make(map[stoneChangeState]int)

type stoneChangeState struct {
	S               stone
	BlinksRemaining int
}

func (d Day11Task1) CalculateAnswer(input string) (string, error) {
	stones := strings.Split(strings.TrimRight(input, "\n"), " ")
	sum := 0
	for _, s := range stones {
		sum += Day11(d).countWithBlinksRemaining(stone(s), 25)
	}
	return strconv.Itoa(sum), nil
}

func (d Day11Task2) CalculateAnswer(input string) (string, error) {
	stones := strings.Split(strings.TrimRight(input, "\n"), " ")
	sum := 0
	for _, s := range stones {
		sum += Day11(d).countWithBlinksRemaining(stone(s), 75)
	}
	return strconv.Itoa(sum), nil
}

func (d Day11) countWithBlinksRemaining(stone stone, blinksRemaining int) int {
	if finalCount, ok := stoneChangeHelper[stoneChangeState{S: stone, BlinksRemaining: blinksRemaining}]; ok {
		return finalCount
	}

	sum := 0

	if blinksRemaining == 0 {
		sum = 1
	} else {
		for _, s := range stone.blink() {
			sum += d.countWithBlinksRemaining(s, blinksRemaining-1)
		}
	}
	stoneChangeHelper[stoneChangeState{S: stone, BlinksRemaining: blinksRemaining}] = sum
	return sum
}

func (s stone) blink() []stone {
	if s == "0" {
		return []stone{"1"}
	}
	if len(s)%2 == 0 {
		half := len(s) / 2
		firstStone := stone(s[0:half])
		secondStone := stone(strings.TrimLeft(string(s[half:]), "0"))
		if secondStone == "" {
			secondStone = "0"
		}
		return []stone{firstStone, secondStone}
	}
	return []stone{stone(strconv.Itoa(s.int() * 2024))}
}

func (s stone) int() int {
	val, _ := strconv.Atoi(string(s))
	return val
}
