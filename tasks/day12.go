package tasks

import (
	"slices"
	"sort"
	"strconv"

	"github.com/nagybalint/advent-of-code-2024/utils"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Day12Task1 struct{}
type Day12Task2 struct{}
type plant string
type garden utils.Plane[plant]
type region sets.Set[utils.Point]

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
				r := garden.getRegionOf(p, visited)
				price += r.area() * r.getBorderingFences(garden)
			}
		}
	}

	return strconv.Itoa(price), nil
}

func (d Day12Task2) CalculateAnswer(input string) (string, error) {
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
				r := garden.getRegionOf(p, visited)
				price += r.area() * r.getSides()
			}
		}
	}

	return strconv.Itoa(price), nil
}

func (r region) getSides() int {
	return r.getTopSides() + r.getBottomSides() + r.getLeftSides() + r.getRightSides()
}

func (r region) getLeftSides() int {
	hLower, hUpper := r.getVerticalBounds()
	leftSides := 0
	for i := hLower; i <= hUpper; i++ {
		var leftCrossSection []utils.Point
		if i > hLower {
			leftCrossSection = r.getVerticalCrossSection(i - 1)
		}
		crossSection := r.getVerticalCrossSection(i)
		exposedOnLeft := collectExposed(crossSection, leftCrossSection, func(p utils.Point) utils.Point {
			return *p.Step(p.XGoesLeft, p.StaysStill)
		})
		leftSides += countExposedSections(exposedOnLeft, func(p utils.Point) int { return p.Y })
	}
	return leftSides
}

func (r region) getRightSides() int {
	hLower, hUpper := r.getVerticalBounds()
	rightSides := 0
	for i := hLower; i <= hUpper; i++ {
		var rightCrossSection []utils.Point
		if i < hUpper {
			rightCrossSection = r.getVerticalCrossSection(i + 1)
		}
		crossSection := r.getVerticalCrossSection(i)
		exposedOnRight := collectExposed(crossSection, rightCrossSection, func(p utils.Point) utils.Point {
			return *p.Step(p.XGoesRight, p.StaysStill)
		})
		rightSides += countExposedSections(exposedOnRight, func(p utils.Point) int { return p.Y })
	}
	return rightSides
}

func (r region) getTopSides() int {
	hLower, hUpper := r.getHorizontalBounds()
	topSides := 0
	for i := hLower; i <= hUpper; i++ {
		var topCrossSection []utils.Point
		if i > hLower {
			topCrossSection = r.getHorizontalCrossSection(i - 1)
		}
		crossSection := r.getHorizontalCrossSection(i)
		exposedOnTop := collectExposed(crossSection, topCrossSection, func(p utils.Point) utils.Point {
			return *p.Step(p.StaysStill, p.YGoesUp)
		})
		topSides += countExposedSections(exposedOnTop, func(p utils.Point) int { return p.X })
	}
	return topSides
}

func (r region) getBottomSides() int {
	hLower, hUpper := r.getHorizontalBounds()
	bottomSides := 0
	for i := hLower; i <= hUpper; i++ {
		var bottomCrossSection []utils.Point
		if i < hUpper {
			bottomCrossSection = r.getHorizontalCrossSection(i + 1)
		}
		crossSection := r.getHorizontalCrossSection(i)
		exposedOnBottom := collectExposed(crossSection, bottomCrossSection, func(p utils.Point) utils.Point {
			return *p.Step(p.StaysStill, p.YGoesDown)
		})
		bottomSides += countExposedSections(exposedOnBottom, func(p utils.Point) int { return p.X })
	}
	return bottomSides
}

func collectExposed(crossSection, checkAgainst []utils.Point, step func(utils.Point) utils.Point) (exposed []utils.Point) {
	for _, p := range crossSection {
		if !slices.Contains(checkAgainst, step(p)) {
			exposed = append(exposed, p)
		}
	}
	return exposed
}

func countExposedSections(exposed []utils.Point, getCoord func(utils.Point) int) int {
	if len(exposed) < 2 {
		return len(exposed)
	}
	sections := 1
	for w := range utils.SlidingWindow(exposed, 2) {
		if getCoord(w[0])+1 != getCoord(w[1]) {
			sections++
		}
	}
	return sections
}

func (r region) getHorizontalBounds() (lower int, upper int) {
	for p := range r {
		if p.Y < lower {
			lower = p.Y
		}
		if p.Y > upper {
			upper = p.Y
		}
	}
	return lower, upper
}

func (r region) getHorizontalCrossSection(y int) (crossSection []utils.Point) {
	for p := range r {
		if p.Y == y {
			crossSection = append(crossSection, p)
		}
	}
	sort.Slice(crossSection, func(i, j int) bool {
		return crossSection[i].X < crossSection[j].X
	})
	return crossSection
}

func (r region) getVerticalBounds() (lower int, upper int) {
	for p := range r {
		if p.X < lower {
			lower = p.X
		}
		if p.X > upper {
			upper = p.X
		}
	}
	return lower, upper
}

func (r region) getVerticalCrossSection(x int) (crossSection []utils.Point) {
	for p := range r {
		if p.X == x {
			crossSection = append(crossSection, p)
		}
	}
	sort.Slice(crossSection, func(i, j int) bool {
		return crossSection[i].Y < crossSection[j].Y
	})
	return crossSection
}

func (r region) area() int {
	return len(r)
}

func (r region) getBorderingFences(g garden) int {
	fences := 0
	for p := range r {
		fences += g.borderingFences(p)
	}
	return fences
}

func (g garden) getRegionOf(p utils.Point, visited utils.Plane[bool]) region {
	r := make(region)
	sets.Set[utils.Point](r).Insert(p)
	visited.SetValueAt(p, true)
	for _, neighbor := range p.Neighbors() {
		if utils.Plane[plant](g).TestValueAt(neighbor, utils.Plane[plant](g).ValueAt(p)) &&
			visited.TestValueAt(neighbor, false) {
			neighborRegion := sets.Set[utils.Point](g.getRegionOf(neighbor, visited))
			r = region(sets.Set[utils.Point](r).Union(neighborRegion))
		}
	}
	return r
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
