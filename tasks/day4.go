package tasks

import (
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day4Task1 struct{}
type Letters [][]rune
type Pos struct {
	x, y int
}

func (Day4Task1) CalculateAnswer(input string) (string, error) {
	lines := utils.Filter(strings.Split(input, "\n"), func(s string) bool {
		return s != ""
	})

	var letters Letters
	for _, l := range lines {
		letters = append(letters, []rune(l))
	}

	xmasCount := 0
	for y := range letters {
		for x := range letters[y] {
			xmasFromPos := 0
			pos := Pos{x, y}
			if letters.hasXmasFrom(pos, xGoesRight, yStaysStill) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xGoesLeft, yStaysStill) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xStaysStill, yGoesDown) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xStaysStill, yGoesUp) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xGoesRight, yGoesUp) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xGoesRight, yGoesDown) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xGoesLeft, yGoesUp) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(pos, xGoesLeft, yGoesDown) {
				xmasFromPos++
			}
			log.Println(pos, " - counted xmases ", xmasFromPos)
			xmasCount += xmasFromPos
		}
	}

	return strconv.Itoa(xmasCount), nil
}

func xGoesRight(x int) int {
	return x + 1
}

func xGoesLeft(x int) int {
	return x - 1
}

func xStaysStill(x int) int {
	return x
}

func yGoesDown(y int) int {
	return y + 1
}

func yGoesUp(y int) int {
	return y - 1
}

func yStaysStill(y int) int {
	return y
}

type Day4Task2 struct{}

func (Day4Task2) CalculateAnswer(input string) (string, error) {
	lines := utils.Filter(strings.Split(input, "\n"), func(s string) bool {
		return s != ""
	})

	var letters Letters
	for _, l := range lines {
		letters = append(letters, []rune(l))
	}
	crossmasCount := 0
	for y := 1; y < len(letters)-1; y++ {
		for x := 1; x < len(letters[y])-1; x++ {
			p := Pos{x, y}
			if letters.hasCrossMasAt(p) {
				crossmasCount++
			}
		}
	}
	return strconv.Itoa(crossmasCount), nil
}

// pos will be used to locate the 'A' of the cross mas
func (l Letters) hasCrossMasAt(pos Pos) bool {
	if pos.x < 1 || pos.y < 1 || pos.y >= len(l)-1 || pos.x >= len(l[pos.y])-1 {
		return false
	}
	if !l.isLetterAt(pos, 'A') {
		return false
	}
	topLeftLetter := l.letterAt(Pos{pos.x - 1, pos.y - 1})
	topRightLetter := l.letterAt(Pos{pos.x - 1, pos.y + 1})
	bottomLeftLetter := l.letterAt(Pos{pos.x + 1, pos.y - 1})
	bottomRightLetter := l.letterAt(Pos{pos.x + 1, pos.y + 1})
	lettersCounts := make(map[rune]int)
	lettersCounts[topLeftLetter]++
	lettersCounts[topRightLetter]++
	lettersCounts[bottomLeftLetter]++
	lettersCounts[bottomRightLetter]++
	if lettersCounts['M'] != 2 {
		return false
	}
	if lettersCounts['S'] != 2 {
		return false
	}
	if topLeftLetter == bottomRightLetter {
		return false
	}
	log.Println("X-Mas found at ", pos)
	return true
}

func (l Letters) letterAt(pos Pos) rune {
	return l[pos.y][pos.x]
}

func (l Letters) hasXmasFrom(pos Pos, xStep, yStep func(int) int) bool {
	if !l.isLetterAt(pos, 'X') {
		return false
	}

	pos.step(xStep, yStep)
	if !l.isLetterAt(pos, 'M') {
		return false
	}

	pos.step(xStep, yStep)
	if !l.isLetterAt(pos, 'A') {
		return false
	}

	pos.step(xStep, yStep)
	if !l.isLetterAt(pos, 'S') {
		return false
	}

	return true
}

func (l Letters) isLetterAt(pos Pos, letter rune) bool {
	if pos.x < 0 {
		return false
	}

	if pos.y < 0 {
		return false
	}

	if pos.y >= len(l) {
		return false
	}

	if pos.x >= len(l[pos.y]) {
		return false
	}

	return l[pos.y][pos.x] == letter
}

func (p *Pos) step(xStep, yStep func(int) int) {
	p.x = xStep(p.x)
	p.y = yStep(p.y)
}
