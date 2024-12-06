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
