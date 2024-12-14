package utils

import (
	"strconv"
	"strings"
)

func BuildPlaneOfRunes(input string) (plane Plane[rune]) {
	for _, line := range Filter(strings.Split(input, "\n"), IsNonEmptyString) {
		var l []rune
		for _, r := range line {
			l = append(l, r)
		}
		plane = append(plane, l)
	}
	return plane
}
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
