package tasks

import (
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day15 struct{}
type Day15Task1 struct{}

const (
	robotBody = '@'
	box       = 'O'
	obstacle  = '#'
	empty     = '.'
)

func (d Day15Task1) CalculateAnswer(input string) (string, error) {
	warehouse, movements := Day15(d).parseInput(input)
	robot := Day15(d).findRobot(warehouse)

	for _, dir := range movements {
		robot.Dir = utils.Direction(dir)
		d.move(&robot, robotBody, warehouse)
	}
	finalBoxes := warehouse.FindAllPointsOfValue(box)
	sum := 0
	for _, b := range finalBoxes {
		sum += d.gpsCoordinate(b)
	}

	log.Println(utils.RunePlaneToString(warehouse))
	return strconv.Itoa(sum), nil
}

func (d Day15Task1) move(walker *utils.Walker, body rune, warehouse utils.Plane[rune]) (hasMoved bool) {
	peekObj, peekPos, peekErr := walker.Peek(warehouse)
	if peekErr != nil {
		log.Fatalln(peekErr)
	}
	switch peekObj {
	case obstacle:
		return false
	case empty:
		warehouse.SetValueAt(walker.Pos, empty)
		warehouse.SetValueAt(peekPos, body)
		walker.Move()
		return true
	case box:
		boxWalker := utils.Walker{Pos: peekPos, Dir: walker.Dir}
		hasBoxMoved := d.move(&boxWalker, box, warehouse)
		if hasBoxMoved {
			warehouse.SetValueAt(walker.Pos, empty)
			warehouse.SetValueAt(peekPos, body)
			walker.Move()
			return true
		} else {
			return false
		}
	}
	return false
}

func (d Day15Task1) gpsCoordinate(pos utils.Point) int {
	return pos.Y*100 + pos.X
}

func (d Day15) parseInput(input string) (warehouse utils.Plane[rune], movements []string) {
	parts := strings.Split(input, "\n\n")
	warehouse = utils.BuildPlaneOfRunes(parts[0])
	for _, line := range utils.Filter(strings.Split(parts[1], "\n"), utils.IsNonEmptyString) {
		for _, r := range line {
			movements = append(movements, string(r))
		}
	}
	return warehouse, movements
}

func (d Day15) findRobot(layout utils.Plane[rune]) utils.Walker {
	return utils.Walker{
		Pos: *layout.FindPointOfValue(robotBody),
		Dir: "",
	}
}
