package utils

type Pos struct {
	X, Y int
}

func (p *Pos) Step(xStep, yStep func(int) int) *Pos {
	return &Pos{xStep(p.X), yStep(p.Y)}
}

func (Pos) GoesRight(x int) int {
	return x + 1
}

func (Pos) GoesLeft(x int) int {
	return x - 1
}

func (Pos) StaysStill(coord int) int {
	return coord
}

func (Pos) GoesDown(y int) int {
	return y + 1
}

func (Pos) GoesUp(y int) int {
	return y - 1
}
