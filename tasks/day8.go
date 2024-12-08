package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day8Task1 struct{}

type antenna struct {
	p   utils.Pos
	typ rune
}

func (a antenna) String() string {
	return fmt.Sprintf("antenna{x: %d, y: %d, typ: %s}", a.p.X, a.p.Y, string(a.typ))
}

type cityLayout [][]rune

func (Day8Task1) CalculateAnswer(input string) (string, error) {
	var city cityLayout
	for _, line := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		var cityLine []rune
		for _, r := range line {
			cityLine = append(cityLine, r)
		}
		city = append(city, cityLine)
	}
	antennas := city.findAntennas()
	grouped := utils.GroupBy(antennas, func(a antenna) rune { return a.typ })
	antinodes := make(sets.Set[utils.Pos])
	for _, antennas := range grouped {
		for pair := range utils.Pairs(antennas) {
			reflection1 := pair[0].p.Reflect(pair[1].p)
			reflection2 := pair[1].p.Reflect(pair[0].p)
			if city.isInBounds(reflection1) {
				antinodes.Insert(reflection1)
			}
			if city.isInBounds(reflection2) {
				antinodes.Insert(reflection2)
			}
		}
	}
	return strconv.Itoa(len(antinodes)), nil
}

type Day8Task2 struct{}

func (Day8Task2) CalculateAnswer(input string) (string, error) {
	var city cityLayout
	for _, line := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		var cityLine []rune
		for _, r := range line {
			cityLine = append(cityLine, r)
		}
		city = append(city, cityLine)
	}
	antennas := city.findAntennas()
	grouped := utils.GroupBy(antennas, func(a antenna) rune { return a.typ })
	antinodes := make(sets.Set[utils.Pos])
	for _, antennas := range grouped {
		for pair := range utils.Pairs(antennas) {
			antinodes.Insert(getAllReflections(city, pair[0].p, pair[1].p)...)
			antinodes.Insert(getAllReflections(city, pair[1].p, pair[0].p)...)
		}
	}
	return strconv.Itoa(len(antinodes)), nil
}

func getAllReflections(city cityLayout, p utils.Pos, center utils.Pos) (reflections []utils.Pos) {
	reflections = append(reflections, center)
	for {
		reflection := p.Reflect(center)
		if city.isInBounds(reflection) {
			reflections = append(reflections, reflection)
			p = center
			center = reflection
		} else {
			break
		}
	}
	return reflections
}

func (layout cityLayout) isInBounds(p utils.Pos) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}
	if p.Y >= len(layout) || p.X >= len(layout[p.Y]) {
		return false
	}
	return true
}

func (layout cityLayout) findAntennas() (antennas []antenna) {
	for y, line := range layout {
		for x, r := range line {
			if r != '.' {
				antennas = append(antennas, antenna{
					p:   utils.Pos{X: x, Y: y},
					typ: r,
				})
			}
		}
	}
	return antennas
}
