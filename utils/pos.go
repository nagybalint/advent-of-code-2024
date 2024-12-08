package utils

type Pos struct {
	X, Y int
}

func (p *Pos) Step(xStep, yStep func(int) int) *Pos {
	return &Pos{xStep(p.X), yStep(p.Y)}
}

func (Pos) XGoesRight(x int) int {
	return x + 1
}

func (Pos) XGoesLeft(x int) int {
	return x - 1
}

func (Pos) StaysStill(coord int) int {
	return coord
}

func (Pos) YGoesDown(y int) int {
	return y + 1
}

func (Pos) YGoesUp(y int) int {
	return y - 1
}

func (p Pos) Reflect(center Pos) Pos {
	return Pos{
		X: p.X + 2*(center.X-p.X),
		Y: p.Y + 2*(center.Y-p.Y),
	}
}
