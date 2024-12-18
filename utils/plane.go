package utils

import "strings"

type Plane[T comparable] [][]T

func (plane Plane[T]) IsInBounds(p Point) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}
	if p.Y >= len(plane) || p.X >= len(plane[p.Y]) {
		return false
	}
	return true
}

func (plane Plane[T]) Clone() Plane[T] {
	var copy Plane[T]
	for _, row := range plane {
		var rowCopy []T
		for _, l := range row {
			rowCopy = append(rowCopy, l)
		}
		copy = append(copy, rowCopy)
	}
	return copy
}

func (plane Plane[T]) ValueAt(p Point) T {
	return plane[p.Y][p.X]
}

func (plane Plane[T]) TestValueAt(p Point, shouldBe T) bool {
	if !plane.IsInBounds(p) {
		return false
	}
	return plane.ValueAt(p) == shouldBe
}

func (plane Plane[T]) SetValueAt(p Point, value T) {
	plane[p.Y][p.X] = value
}

func (plane Plane[T]) FindPointOfValue(value T) (p *Point) {
	for y, row := range plane {
		for x, v := range row {
			if v == value {
				return &Point{X: x, Y: y}
			}
		}
	}
	return nil
}

func (plane Plane[T]) FindAllPointsOfValue(value T) (positions []Point) {
	for y, row := range plane {
		for x, v := range row {
			if v == value {
				positions = append(positions, Point{X: x, Y: y})
			}
		}
	}
	return positions
}

func MapPlane[T comparable, U comparable](plane Plane[T], mapF func(val T) U) (mapped Plane[U]) {
	for _, row := range plane {
		var mappedRow []U
		for _, v := range row {
			mappedRow = append(mappedRow, mapF(v))
		}
		mapped = append(mapped, mappedRow)
	}
	return mapped
}

func RunePlaneToString(plane Plane[rune]) string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, line := range plane {
		for _, elem := range line {
			sb.WriteRune(rune(elem))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
