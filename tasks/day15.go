package tasks

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day15 struct{}
type Day15Task1 struct{}
type Day15Task2 struct{}

const (
	robotBody    = '@'
	box          = 'O'
	obstacle     = '#'
	empty        = '.'
	wideBoxLeft  = '['
	wideBoxRight = ']'
)

type wideWalker struct {
	leftWalker, rightWalker *utils.Walker
}

func (ww *wideWalker) setDir(d utils.Direction) {
	ww.leftWalker.Dir = d
	ww.rightWalker.Dir = d
}

func (ww *wideWalker) move() {
	ww.leftWalker.Move()
	ww.rightWalker.Move()
}

func (ww wideWalker) canMove(warehouse utils.Plane[rune]) bool {
	leftPeekObj, _, _ := ww.leftWalker.Peek(warehouse)
	rightPeekObj, _, _ := ww.rightWalker.Peek(warehouse)
	nextBoxes := ww.getNextWideBoxes(warehouse)
	switch ww.leftWalker.Dir {
	case utils.Left:
		if leftPeekObj == empty {
			return true
		}
		if leftPeekObj == obstacle {
			return false
		}
		return nextBoxes[0].canMove(warehouse)
	case utils.Right:
		if rightPeekObj == empty {
			return true
		}
		if rightPeekObj == obstacle {
			return false
		}
		return nextBoxes[0].canMove(warehouse)
	case utils.Down, utils.Up:
		if leftPeekObj == empty && rightPeekObj == empty {
			return true
		}
		if leftPeekObj == obstacle || rightPeekObj == obstacle {
			return false
		}
		for _, b := range nextBoxes {
			if !b.canMove(warehouse) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Errorf("Invalid direction:%s", ww.leftWalker.Dir))
	}
}

func (ww *wideWalker) moveInWarehouse(warehouse utils.Plane[rune]) {
	for _, b := range ww.getNextWideBoxes(warehouse) {
		b.moveInWarehouse(warehouse)
	}
	warehouse.SetValueAt(ww.leftWalker.Pos, empty)
	warehouse.SetValueAt(ww.rightWalker.Pos, empty)
	ww.move()
	warehouse.SetValueAt(ww.leftWalker.Pos, wideBoxLeft)
	warehouse.SetValueAt(ww.rightWalker.Pos, wideBoxRight)
}

func (ww wideWalker) getNextWideBoxes(warehouse utils.Plane[rune]) (boxes []wideWalker) {
	if ww.leftWalker.Dir == utils.Left {
		peekObj, peekPos, _ := ww.leftWalker.Peek(warehouse)
		if peekObj == wideBoxRight {
			b := getWideBoxFrom(peekObj, peekPos, ww.leftWalker.Dir)
			if b != nil {
				boxes = append(boxes, *b)
			}
		}
		return boxes
	}
	if ww.leftWalker.Dir == utils.Right {
		peekObj, peekPos, _ := ww.rightWalker.Peek(warehouse)
		if peekObj == wideBoxLeft {
			b := getWideBoxFrom(peekObj, peekPos, ww.leftWalker.Dir)
			if b != nil {
				boxes = append(boxes, *b)
			}
		}
		return boxes
	}
	leftPeekObj, leftPeekPos, _ := ww.leftWalker.Peek(warehouse)
	rightPeekObj, rightPeekPos, _ := ww.rightWalker.Peek(warehouse)
	if leftPeekObj == wideBoxLeft && rightPeekObj == wideBoxRight {
		b := getWideBoxFrom(leftPeekObj, leftPeekPos, ww.leftWalker.Dir)
		if b != nil {
			boxes = append(boxes, *b)
		}
		return boxes
	}
	if leftBox := getWideBoxFrom(leftPeekObj, leftPeekPos, ww.leftWalker.Dir); leftBox != nil {
		boxes = append(boxes, *leftBox)
	}
	if rightBox := getWideBoxFrom(rightPeekObj, rightPeekPos, ww.leftWalker.Dir); rightBox != nil {
		boxes = append(boxes, *rightBox)
	}
	return boxes
}

func getWideBoxFrom(r rune, p utils.Point, d utils.Direction) *wideWalker {
	if r == wideBoxLeft {
		return &wideWalker{
			leftWalker:  &utils.Walker{Pos: p, Dir: d},
			rightWalker: &utils.Walker{Pos: *p.Step(p.XGoesRight, p.StaysStill), Dir: d},
		}
	} else if r == wideBoxRight {
		return &wideWalker{
			leftWalker:  &utils.Walker{Pos: *p.Step(p.XGoesLeft, p.StaysStill), Dir: d},
			rightWalker: &utils.Walker{Pos: p, Dir: d},
		}
	} else {
		return nil
	}
}

func (d Day15Task2) CalculateAnswer(input string) (string, error) {
	warehouse, movements := Day15(d).parseInput(input)
	wideWarehouse := d.wideWarehouse(warehouse)
	log.Println(utils.RunePlaneToString(wideWarehouse))
	robot := Day15(d).findRobot(wideWarehouse)

	for _, dir := range movements {
		robot.Dir = utils.Direction(dir)
		d.move(&robot, wideWarehouse)
	}

	log.Println(utils.RunePlaneToString(wideWarehouse))

	sum := 0
	for _, b := range wideWarehouse.FindAllPointsOfValue(wideBoxLeft) {
		sum += Day15(d).gpsCoordinate(b)
	}

	return strconv.Itoa(sum), nil
}

func (d Day15Task2) move(walker *utils.Walker, warehouse utils.Plane[rune]) {
	peekObj, peekPos, peekErr := walker.Peek(warehouse)
	if peekErr != nil {
		log.Fatalln(peekErr)
	}
	if peekObj == obstacle {
		return
	}
	if peekObj == empty {
		warehouse.SetValueAt(walker.Pos, empty)
		warehouse.SetValueAt(peekPos, robotBody)
		walker.Move()
		return
	}
	nextBox := getWideBoxFrom(peekObj, peekPos, walker.Dir)
	if nextBox.canMove(warehouse) {
		nextBox.moveInWarehouse(warehouse)
		warehouse.SetValueAt(walker.Pos, empty)
		warehouse.SetValueAt(peekPos, robotBody)
		walker.Move()
	}
	return
}

func (d Day15Task2) wideWarehouse(warehouse utils.Plane[rune]) utils.Plane[rune] {
	var wideWarehouse utils.Plane[rune]
	for _, row := range warehouse {
		var wideRow []rune
		for _, elem := range row {
			switch elem {
			case empty:
				wideRow = append(wideRow, empty, empty)
			case obstacle:
				wideRow = append(wideRow, obstacle, obstacle)
			case robotBody:
				wideRow = append(wideRow, robotBody, empty)
			case box:
				wideRow = append(wideRow, wideBoxLeft, wideBoxRight)
			default:
				panic(fmt.Errorf("Unexpected element while widening map"))
			}
		}
		wideWarehouse = append(wideWarehouse, wideRow)
	}
	return wideWarehouse
}

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
		sum += Day15(d).gpsCoordinate(b)
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

func (d Day15) gpsCoordinate(pos utils.Point) int {
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
