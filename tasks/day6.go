package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day6Task1 struct{}

type guard struct {
	pos utils.Pos
	or  orientation
}
type orientation string

const (
	up    orientation = "^"
	down  orientation = "v"
	left  orientation = "<"
	right orientation = ">"
)

type labLayout [][]rune

func (Day6Task1) CalculateAnswer(input string) (string, error) {
	var layout labLayout
	for _, l := range utils.Filter(strings.Split(input, "\n"), func(s string) bool { return s != "" }) {
		var line []rune
		for _, r := range l {
			line = append(line, r)
		}
		layout = append(layout, line)
	}

	g := findGuard(layout)
	for g.isInBounds(layout) {
		if layout[g.pos.Y][g.pos.X] == '#' {
			g.stepBack()
			g.turn()
		} else {
			layout[g.pos.Y][g.pos.X] = 'X'
			g.move()
		}
	}

	var visited int
	for _, l := range layout {
		for _, r := range l {
			if r == 'X' {
				visited++
			}
		}
	}

	return strconv.Itoa(visited), nil
}

func findGuard(layout labLayout) *guard {
	for y, line := range layout {
		for x, r := range line {
			if r != '#' && r != '.' {
				return &guard{
					pos: utils.Pos{X: x, Y: y},
					or:  orientation(r),
				}
			}
		}
	}
	return nil
}

func (g *guard) move() {
	switch g.or {
	case up:
		g.pos = *g.pos.Step(g.pos.StaysStill, g.pos.GoesUp)
	case down:
		g.pos = *g.pos.Step(g.pos.StaysStill, g.pos.GoesDown)
	case left:
		g.pos = *g.pos.Step(g.pos.GoesLeft, g.pos.StaysStill)
	case right:
		g.pos = *g.pos.Step(g.pos.GoesRight, g.pos.StaysStill)
	default:
		panic(fmt.Errorf("Invalid orientation"))
	}
}

func (g *guard) stepBack() {
	g.or = g.or.opposite()
	g.move()
	g.or = g.or.opposite()
}

func (g *guard) turn() {
	o := &g.or
	switch *o {
	case up:
		*o = right
	case down:
		*o = left
	case left:
		*o = up
	case right:
		*o = down
	default:
		panic(fmt.Errorf("Invalid orientation"))
	}
}

func (g *guard) isInBounds(layout labLayout) bool {
	if g.pos.X < 0 {
		return false
	}
	if g.pos.Y < 0 {
		return false
	}
	if g.pos.Y >= len(layout) {
		return false
	}
	if g.pos.X >= len(layout[g.pos.Y]) {
		return false
	}
	return true
}

func (o orientation) opposite() orientation {
	switch o {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	default:
		panic(fmt.Errorf("Invalid orientation"))
	}
}
