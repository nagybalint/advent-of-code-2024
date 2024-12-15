package tasks

import (
	"strconv"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day6Task1 struct{}

type visits map[utils.Point]sets.Set[utils.Direction]

func (Day6Task1) CalculateAnswer(input string) (string, error) {
	layout := utils.BuildPlaneOfRunes(input)

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
	layout := utils.BuildPlaneOfRunes(input)

	visits := make(visits)
	newObstacles := make(sets.Set[utils.Point])
	g := findGuard(layout)
	for layout.IsInBounds(g.Pos) {
		if layout[g.Pos.Y][g.Pos.X] == '#' {
			g.StepBack()
			g.TurnRight()
		} else {
			nextField, nextPos, err := g.Peek(layout)
			// Check if we would still be inbounds if we moved
			if err == nil {
				_, wasNextPosVisited := visits[nextPos]
				// If there is an obstacle on the next forward step, or
				// we already visited the field that we would use to test a new obstacle, don't add the obstacle
				if nextField != '#' && !wasNextPosVisited {
					testLayout := layout.Clone()
					testLayout.SetValueAt(nextPos, '#')
					_, addingOstacleResultsInLoop := completePatrol(*g, testLayout, visits)
					if addingOstacleResultsInLoop {
						newObstacles.Insert(nextPos)
					}
				}
			}
			visits.add(g.Pos, g.Dir)
			g.Move()
		}
	}

	return strconv.Itoa(len(newObstacles)), nil
}

func completePatrol(g utils.Walker, layout utils.Plane[rune], alreadyVisited visits) (visits visits, hasLoop bool) {
	visits = alreadyVisited.clone()
	for {
		if !layout.IsInBounds(g.Pos) {
			return visits, false
		}
		if visits.has(g.Pos, g.Dir) {
			return visits, true
		}
		if layout.TestValueAt(g.Pos, '#') {
			g.StepBack()
			g.TurnRight()
		} else {
			visits.add(g.Pos, g.Dir)
			g.Move()
		}
	}
}

func (v visits) add(p utils.Point, o utils.Direction) {
	if _, ok := v[p]; !ok {
		v[p] = make(sets.Set[utils.Direction])
	}
	v[p].Insert(o)
}

func (v visits) has(p utils.Point, o utils.Direction) bool {
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

func findGuard(layout utils.Plane[rune]) *utils.Walker {
	for y, line := range layout {
		for x, r := range line {
			if r != '#' && r != '.' {
				return &utils.Walker{
					Pos: utils.Point{X: x, Y: y},
					Dir: utils.Direction(r),
				}
			}
		}
	}
	return nil
}
