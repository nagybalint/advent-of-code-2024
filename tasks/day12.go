package tasks

import (
	"strconv"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day12Task1 struct{}
type plant string
type garden utils.Plane[plant]

func (d Day12Task1) CalculateAnswer(input string) (string, error) {
	garden := garden(utils.MapPlane(utils.BuildPlaneOfRunes(input), func(s rune) plant { return plant(s) }))
	var visited utils.Plane[bool] = make([][]bool, len(garden))
	for x := range garden {
		visited[x] = make([]bool, len(garden[x]))
	}
	price := 0
	for y := range garden {
		for x := range garden[y] {
			p := utils.Point{X: x, Y: y}
			if visited.TestValueAt(p, false) {
				region := garden.getRegionOf(p, visited)
				fences := 0
				for _p := range region {
					fences += garden.borderingFences(_p)
				}
				price += len(region) * fences
			}
		}
	}

	return strconv.Itoa(price), nil
}

func (g garden) getRegionOf(p utils.Point, visited utils.Plane[bool]) sets.Set[utils.Point] {
	region := make(sets.Set[utils.Point])
	region.Insert(p)
	visited.SetValueAt(p, true)
	for _, neighbor := range p.Neighbors() {
		if utils.Plane[plant](g).TestValueAt(neighbor, utils.Plane[plant](g).ValueAt(p)) &&
			visited.TestValueAt(neighbor, false) {
			region = region.Union(g.getRegionOf(neighbor, visited))
		}
	}
	return region
}

func (g garden) borderingFences(p utils.Point) int {
	fences := 0
	for _, neighbor := range p.Neighbors() {
		if !utils.Plane[plant](g).TestValueAt(neighbor, utils.Plane[plant](g).ValueAt(p)) {
			fences += 1
		}
	}
	return fences
}
