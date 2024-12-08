package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day6Task1 struct{}

type guard struct {
	pos utils.Point
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

type visits map[utils.Point]sets.Set[orientation]

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
	emptyVisits := make(visits)
	visits, _ := completePatrol(*g, layout, emptyVisits)

	var visited int
	for x, _ := range visits {
		visited += len(visits[x])
	}

	return strconv.Itoa(visited), nil
}

type Day6Task2 struct{}

func (Day6Task2) CalculateAnswer(input string) (string, error) {
	var layout labLayout
	for _, l := range utils.Filter(strings.Split(input, "\n"), func(s string) bool { return s != "" }) {
		var line []rune
		for _, r := range l {
			line = append(line, r)
		}
		layout = append(layout, line)
	}

	visits := make(visits)
	newObstacles := make(sets.Set[utils.Point])
	g := findGuard(layout)
	for g.isInBounds(layout) {
		if layout[g.pos.Y][g.pos.X] == '#' {
			g.stepBack()
			g.turn()
		} else {
			nextField, nextPos, err := g.peek(layout)
			// Check if we would still be inbounds if we moved
			if err == nil {
				_, wasNextPosVisited := visits[nextPos]
				// If there is an obstacle on the next forward step, or
				// we already visited the field that we would use to test a new obstacle, don't add the obstacle
				if nextField != '#' && !wasNextPosVisited {
					testLayout := layout.clone()
					testLayout[nextPos.Y][nextPos.X] = '#'
					_, addingOstacleResultsInLoop := completePatrol(*g, testLayout, visits)
					if addingOstacleResultsInLoop {
						newObstacles.Insert(nextPos)
					}
				}
			}
			visits.add(g.pos, g.or)
			g.move()
		}
	}

	return strconv.Itoa(len(newObstacles)), nil
}

func completePatrol(g guard, layout labLayout, alreadyVisited visits) (visits visits, hasLoop bool) {
	visits = alreadyVisited.clone()
	for {
		if !g.isInBounds(layout) {
			return visits, false
		}
		if visits.has(g.pos, g.or) {
			return visits, true
		}
		if layout[g.pos.Y][g.pos.X] == '#' {
			g.stepBack()
			g.turn()
		} else {
			visits.add(g.pos, g.or)
			g.move()
		}
	}
}

func (g guard) isNextStepObstacleable(layout labLayout, v visits) bool {
	g.turn()
	for g.isInBounds(layout) && layout[g.pos.Y][g.pos.X] != '#' {
		if v.has(g.pos, g.or) {
			return true
		}
	}
	return false
}

func (v visits) add(p utils.Point, o orientation) {
	if _, ok := v[p]; !ok {
		v[p] = make(sets.Set[orientation])
	}
	v[p].Insert(o)
}

func (v visits) has(p utils.Point, o orientation) bool {
	if _, ok := v[p]; !ok {
		return false
	}
	return v[p].Has(o)
}

func (v visits) clone() visits {
	copy := make(visits)
	for p := range v {
		copy[p] = v[p].Clone()
	}
	return copy
}

func findGuard(layout labLayout) *guard {
	for y, line := range layout {
		for x, r := range line {
			if r != '#' && r != '.' {
				return &guard{
					pos: utils.Point{X: x, Y: y},
					or:  orientation(r),
				}
			}
		}
	}
	return nil
}

func (l labLayout) clone() labLayout {
	var copy labLayout
	for _, line := range l {
		var lineCopy []rune
		for _, l := range line {
			lineCopy = append(lineCopy, l)
		}
		copy = append(copy, lineCopy)
	}
	return copy
}

func (g *guard) move() {
	switch g.or {
	case up:
		g.pos = *g.pos.Step(g.pos.StaysStill, g.pos.YGoesUp)
	case down:
		g.pos = *g.pos.Step(g.pos.StaysStill, g.pos.YGoesDown)
	case left:
		g.pos = *g.pos.Step(g.pos.XGoesLeft, g.pos.StaysStill)
	case right:
		g.pos = *g.pos.Step(g.pos.XGoesRight, g.pos.StaysStill)
	default:
		panic(fmt.Errorf("Invalid orientation"))
	}
}

func (g guard) peek(layout labLayout) (elem rune, pos utils.Point, err error) {
	g.move()
	if !g.isInBounds(layout) {
		err = fmt.Errorf("Guard out of bounds")
		return
	}
	elem = layout[g.pos.Y][g.pos.X]
	pos = g.pos
	return
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

func (g *guard) turnBack() {
	g.or = g.or.opposite()
	g.turn()
}

func (g *guard) isInBounds(layout labLayout) bool {
	return layout.isInBounds(g.pos)
}

func (layout labLayout) isInBounds(p utils.Point) bool {
	if p.X < 0 {
		return false
	}
	if p.Y < 0 {
		return false
	}
	if p.Y >= len(layout) {
		return false
	}
	if p.X >= len(layout[p.Y]) {
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
