package tasks

import (
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day4Task1 struct{}
type Letters utils.Plane[rune]

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
			pos := utils.Point{X: x, Y: y}
			if letters.hasXmasFrom(&pos, pos.XGoesRight, pos.StaysStill) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.XGoesLeft, pos.StaysStill) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.StaysStill, pos.YGoesDown) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.StaysStill, pos.YGoesUp) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.XGoesRight, pos.StaysStill) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.XGoesRight, pos.StaysStill) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.XGoesLeft, pos.YGoesUp) {
				xmasFromPos++
			}
			if letters.hasXmasFrom(&pos, pos.XGoesLeft, pos.YGoesDown) {
				xmasFromPos++
			}
			log.Println(pos, " - counted xmases ", xmasFromPos)
			xmasCount += xmasFromPos
		}
	}

	return strconv.Itoa(xmasCount), nil
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
			p := utils.Point{X: x, Y: y}
			if letters.hasCrossMasAt(p) {
				crossmasCount++
			}
		}
	}
	return strconv.Itoa(crossmasCount), nil
}

// pos will be used to locate the 'A' of the cross mas
func (l Letters) hasCrossMasAt(pos utils.Point) bool {
	if pos.X < 1 || pos.Y < 1 || pos.Y >= len(l)-1 || pos.X >= len(l[pos.Y])-1 {
		return false
	}
	if !utils.Plane[rune](l).TestValueAt(pos, 'A') {
		return false
	}
	topLeftLetter := utils.Plane[rune](l).ValueAt(*pos.Step(pos.XGoesLeft, pos.YGoesUp))
	topRightLetter := utils.Plane[rune](l).ValueAt(*pos.Step(pos.XGoesRight, pos.YGoesUp))
	bottomLeftLetter := utils.Plane[rune](l).ValueAt(*pos.Step(pos.XGoesLeft, pos.YGoesDown))
	bottomRightLetter := utils.Plane[rune](l).ValueAt(*pos.Step(pos.XGoesRight, pos.YGoesDown))
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

func (l Letters) hasXmasFrom(pos *utils.Point, xStep, yStep func(int) int) bool {
	if !utils.Plane[rune](l).TestValueAt(*pos, 'X') {
		return false
	}

	pos = pos.Step(xStep, yStep)
	if !utils.Plane[rune](l).TestValueAt(*pos, 'M') {
		return false
	}

	pos = pos.Step(xStep, yStep)
	if !utils.Plane[rune](l).TestValueAt(*pos, 'A') {
		return false
	}

	pos = pos.Step(xStep, yStep)
	if !utils.Plane[rune](l).TestValueAt(*pos, 'S') {
		return false
	}

	return true
}
