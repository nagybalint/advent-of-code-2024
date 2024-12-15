package utils

import (
	"fmt"
)

type Walker struct {
	Pos Point
	Dir Direction
}
type Direction string

const (
	Up    Direction = "^"
	Down  Direction = "v"
	Left  Direction = "<"
	Right Direction = ">"
)

func (w *Walker) Move() {
	switch w.Dir {
	case Up:
		w.Pos = *w.Pos.Step(w.Pos.StaysStill, w.Pos.YGoesUp)
	case Down:
		w.Pos = *w.Pos.Step(w.Pos.StaysStill, w.Pos.YGoesDown)
	case Left:
		w.Pos = *w.Pos.Step(w.Pos.XGoesLeft, w.Pos.StaysStill)
	case Right:
		w.Pos = *w.Pos.Step(w.Pos.XGoesRight, w.Pos.StaysStill)
	default:
		panic(fmt.Errorf("Invalid direction: ", w.Dir))
	}
}

func (w Walker) Peek(layout Plane[rune]) (elem rune, pos Point, err error) {
	w.Move()
	if !layout.IsInBounds(w.Pos) {
		err = fmt.Errorf("Walker out of bounds at: ", w.Pos, w.Dir)
		return
	}
	elem = layout.ValueAt(w.Pos)
	pos = w.Pos
	return
}

func (w *Walker) StepBack() {
	w.Dir = w.Dir.Opposite()
	w.Move()
	w.Dir = w.Dir.Opposite()
}

func (w *Walker) TurnRight() {
	d := &w.Dir
	switch *d {
	case Up:
		*d = Right
	case Down:
		*d = Left
	case Left:
		*d = Up
	case Right:
		*d = Down
	default:
		panic(fmt.Errorf("Invalid direction: ", d))
	}
}

func (w *Walker) TurnLeft() {
	w.Dir = w.Dir.Opposite()
	w.TurnRight()
}

func (w *Walker) TurnBack() {
	w.Dir = w.Dir.Opposite()
}

func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		panic(fmt.Errorf("Invalid direction: ", d))
	}
}
