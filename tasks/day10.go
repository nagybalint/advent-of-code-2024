package tasks

import (
	"log"
	"strconv"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day10 struct {}

type Day10Task1 struct{}
type topomap utils.Plane[int]

func (d Day10Task1) CalculateAnswer(input string) (string, error) {
	topo := Day10(d).topo(input)
	zeroPoints := utils.Plane[int](topo).FindAllPointsOfValue(0)

	sum := 0
	for _, zp := range zeroPoints {
		sum += topo.trailheadScoreOf(zp)
	}

	return strconv.Itoa(sum), nil
}

type Day10Task2 struct{}

func (d Day10Task2) CalculateAnswer(input string) (string, error) {
	topo := Day10(d).topo(input)
	zeroPoints := utils.Plane[int](topo).FindAllPointsOfValue(0)

	sum := 0
	for _, zp := range zeroPoints {
		sum += topo.rating(zp)
	}

	return strconv.Itoa(sum), nil
}

func (t topomap) rating(p utils.Point) int {
	if utils.Plane[int](t).TestValueAt(p, 9) {
		return 1
	}
	cms := t.climbOptionsFrom(p)
	if len(cms) == 0 {
		return 0
	}
	r := 0
	for cm := range cms {
		r += t.rating(cm)
	}
	return r
}

func (t topomap) trailheadScoreOf(p utils.Point) int {
	score := 0
	cms := t.climbableMaximumsFrom(p)
	for cm := range cms {
		if utils.Plane[int](t).TestValueAt(cm, 9) {
			score++
		}
	}
	return score
}

func (t topomap) climbableMaximumsFrom(p utils.Point) sets.Set[utils.Point] {
	maximums := make(sets.Set[utils.Point])
	maxOptions := make(sets.Set[utils.Point])
	maxOptions.Insert(p)
	for len(maxOptions) != 0 {
		log.Println(maxOptions)
		for m := range maxOptions.Clone() {
			cos := t.climbOptionsFrom(m)
			maxOptions.Delete(m)
			if len(cos) > 0 {
				maxOptions = maxOptions.Union(cos)
			} else {
				maximums.Insert(m)
			}
		}
	}
	return maximums
}

func (t topomap) climbOptionsFrom(p utils.Point) sets.Set[utils.Point] {
	options := make(sets.Set[utils.Point])
	for _, opt := range []utils.Point{
		*p.Step(p.XGoesLeft, p.StaysStill),
		*p.Step(p.XGoesRight, p.StaysStill),
		*p.Step(p.StaysStill, p.YGoesDown),
		*p.Step(p.StaysStill, p.YGoesUp),
	} {
		if utils.Plane[int](t).IsInBounds(opt) &&
			(utils.Plane[int](t).ValueAt(opt) == utils.Plane[int](t).ValueAt(p)+1) {
			options.Insert(opt)
		}
	}
	return options
}

func (d Day10) topo(input string) topomap {
	var topo utils.Plane[int]
	for _, line := range utils.BuildPlaneOfRunes(input) {
		var topoline []int
		for _, r := range line {
			h, _ := strconv.Atoi(string(r))
			topoline = append(topoline, h)
		}
		topo = append(topo, topoline)
	}
	return topomap(topo)
}
