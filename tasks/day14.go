package tasks

import (
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day14Task1 struct{}

// call ((X1 + V1*N)%W+W)%W

type velocity struct {
	X, Y int
}
type robot struct {
	p utils.Point
	v velocity
}

func (d Day14Task1) CalculateAnswer(input string) (string, error) {
	var robots []robot
	for _, rawRobot := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		robots = append(robots, d.parseRobot(rawRobot))
	}
	steps := 100
	tall := 103
	wide := 101
	//tall := 7
	//wide := 11
	var endPositions []utils.Point
	for _, r := range robots {
		endPositions = append(endPositions, r.getPositionAfter(steps, tall, wide))
	}
	d.visualise(tall, wide, endPositions)
	var quadrants [][]utils.Point
	for _, filterFun := range []func(utils.Point) bool{
		func(p utils.Point) bool { return p.X < wide/2 && p.Y < tall/2 },
		func(p utils.Point) bool { return p.X < wide/2 && p.Y > tall/2 },
		func(p utils.Point) bool { return p.X > wide/2 && p.Y < tall/2 },
		func(p utils.Point) bool { return p.X > wide/2 && p.Y > tall/2 },
	} {
		quadrants = append(quadrants, utils.Filter(endPositions, filterFun))
	}
	acc := 1
	for _, q := range quadrants {
		acc *= len(q)
	}
	return strconv.Itoa(acc), nil
}

func (d Day14Task1) visualise(tall, wide int, pos []utils.Point) {
	plane := make([][]int, tall)
	for y := range plane {
		plane[y] = make([]int, wide)
	}
	for y := range plane {
		for x := range plane[y] {
			plane[y][x] = 0
		}
	}
	for _, p := range pos {
		plane[p.Y][p.X]++
	}
	for y := range plane {
		var line []string
		for x := range plane[y] {
			if plane[y][x] == 0 {
				line = append(line, ".")
			} else {
				line = append(line, strconv.Itoa(plane[y][x]))
			}
		}
		log.Println(line)
	}
	log.Println()
}

func (r robot) getPositionAfter(steps, h, w int) (endPosition utils.Point) {
	return utils.Point{
		X: ((r.p.X+r.v.X*steps)%w + w) % w,
		Y: ((r.p.Y+r.v.Y*steps)%h + h) % h,
	}
}

func (d Day14Task1) parseRobot(raw string) robot {
	parts := strings.Split(raw, " ")
	rawPos := strings.Split(parts[0][2:], ",")
	rawVelicity := strings.Split(parts[1][2:], ",")
	return robot{
		p: utils.Point{X: utils.Atoi(rawPos[0]), Y: utils.Atoi(rawPos[1])},
		v: velocity{X: utils.Atoi(rawVelicity[0]), Y: utils.Atoi(rawVelicity[1])},
	}
}
