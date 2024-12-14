package tasks

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day14 struct{}
type Day14Task1 struct{}

// INTERACTIVE!!!
type Day14Task2 struct{}

type velocity struct {
	X, Y int
}
type robot struct {
	p *utils.Point
	v velocity
}

func (d Day14Task1) CalculateAnswer(input string) (string, error) {
	var robots []robot
	for _, rawRobot := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		robots = append(robots, Day14(d).parseRobot(rawRobot))
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
	Day14(d).visualise(0, tall, wide, endPositions, ".")
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

// Press x to stop iteration; do this if you see the christmas tree
// Press any other button to procced
func (d Day14Task2) CalculateAnswer(input string) (string, error) {
	var robots []robot
	for _, rawRobot := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		robots = append(robots, Day14(d).parseRobot(rawRobot))
	}
	tall := 103
	wide := 101

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// print log statements
	log.SetOutput(os.Stdout)

	var b []byte = make([]byte, 1)
	iterations := 0
	topClusteringFactor := 0
	for {
		positions := make(sets.Set[utils.Point])
		for _, r := range robots {
			positions.Insert(*r.p)
		}
		clusteringFactor := 0
		for p := range positions {
			if positions.HasAny(p.Neighbors()...) {
				clusteringFactor++
			}
		}
		if clusteringFactor >= topClusteringFactor {
			topClusteringFactor = clusteringFactor
			Day14(d).visualise(iterations, tall, wide, positions.UnsortedList(), " ")
			os.Stdin.Read(b)
			if string(b) == "x" {
				break
			}
		}
		for _, r := range robots {
			r.move(tall, wide)
		}
		iterations++
	}

	return strconv.Itoa(iterations), nil
}

func (d Day14) visualise(itertion, tall, wide int, pos []utils.Point, placeholder string) {
	// Clear the console
	log.Print("\033[H\033[2J")

	// Print the visualisation
	log.Println("Iteration ", itertion)
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
				line = append(line, placeholder)
			} else {
				line = append(line, strconv.Itoa(plane[y][x]))
			}
		}
		log.Println(line)
	}
}

func (r robot) getPositionAfter(steps, h, w int) (endPosition utils.Point) {
	return utils.Point{
		X: ((r.p.X+r.v.X*steps)%w + w) % w,
		Y: ((r.p.Y+r.v.Y*steps)%h + h) % h,
	}
}

func (r *robot) move(h, w int) {
	r.p.X = ((r.p.X+r.v.X)%w + w) % w
	r.p.Y = ((r.p.Y+r.v.Y)%h + h) % h
}

func (d Day14) parseRobot(raw string) robot {
	parts := strings.Split(raw, " ")
	rawPos := strings.Split(parts[0][2:], ",")
	rawVelicity := strings.Split(parts[1][2:], ",")
	return robot{
		p: &utils.Point{X: utils.Atoi(rawPos[0]), Y: utils.Atoi(rawPos[1])},
		v: velocity{X: utils.Atoi(rawVelicity[0]), Y: utils.Atoi(rawVelicity[1])},
	}
}
