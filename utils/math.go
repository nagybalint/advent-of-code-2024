package utils

func AbsInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Sign(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}
