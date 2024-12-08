package tasks

import (
	"fmt"
	"strconv"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day8Task1 struct{}

type antenna struct {
	p   utils.Point
	typ rune
}

func (a antenna) String() string {
	return fmt.Sprintf("antenna{x: %d, y: %d, typ: %s}", a.p.X, a.p.Y, string(a.typ))
}

type cityLayout utils.Plane[rune]

func (Day8Task1) CalculateAnswer(input string) (string, error) {
	city := cityLayout(utils.BuildPlaneOfRunes(input))
	antennas := city.findAntennas()
	grouped := utils.GroupBy(antennas, func(a antenna) rune { return a.typ })
	antinodes := make(sets.Set[utils.Point])
	for _, antennas := range grouped {
		for pair := range utils.Pairs(antennas) {
			reflection1 := pair[0].p.Reflect(pair[1].p)
			reflection2 := pair[1].p.Reflect(pair[0].p)
			if utils.Plane[rune](city).IsInBounds(reflection1) {
				antinodes.Insert(reflection1)
			}
			if utils.Plane[rune](city).IsInBounds(reflection2) {
				antinodes.Insert(reflection2)
			}
		}
	}
	return strconv.Itoa(len(antinodes)), nil
}

type Day8Task2 struct{}

func (Day8Task2) CalculateAnswer(input string) (string, error) {
	city := cityLayout(utils.BuildPlaneOfRunes(input))
	antennas := city.findAntennas()
	grouped := utils.GroupBy(antennas, func(a antenna) rune { return a.typ })
	antinodes := make(sets.Set[utils.Point])
	for _, antennas := range grouped {
		for pair := range utils.Pairs(antennas) {
			antinodes.Insert(getAllReflections(city, pair[0].p, pair[1].p)...)
			antinodes.Insert(getAllReflections(city, pair[1].p, pair[0].p)...)
		}
	}
	return strconv.Itoa(len(antinodes)), nil
}

func getAllReflections(city cityLayout, p utils.Point, center utils.Point) (reflections []utils.Point) {
	reflections = append(reflections, center)
	for {
		reflection := p.Reflect(center)
		if utils.Plane[rune](city).IsInBounds(reflection) {
			reflections = append(reflections, reflection)
			p = center
			center = reflection
		} else {
			break
		}
	}
	return reflections
}

func (layout cityLayout) findAntennas() (antennas []antenna) {
	for y, line := range layout {
		for x, r := range line {
			if r != '.' {
				antennas = append(antennas, antenna{
					p:   utils.Point{X: x, Y: y},
					typ: r,
				})
			}
		}
	}
	return antennas
}
