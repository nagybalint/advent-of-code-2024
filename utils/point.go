package utils

type Point struct {
	X, Y int
}

func (p *Point) Step(xStep, yStep func(int) int) *Point {
	return &Point{xStep(p.X), yStep(p.Y)}
}

func (Point) XGoesRight(x int) int {
	return x + 1
}

func (Point) XGoesLeft(x int) int {
	return x - 1
}

func (Point) StaysStill(coord int) int {
	return coord
}

func (Point) YGoesDown(y int) int {
	return y + 1
}

func (Point) YGoesUp(y int) int {
	return y - 1
}

func (p Point) Reflect(center Point) Point {
	return Point{
		X: p.X + 2*(center.X-p.X),
		Y: p.Y + 2*(center.Y-p.Y),
	}
}
