package tasks

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day13Task1 struct {
	// If we generalize the input to the following
	// Button A: X+o, Y+p
	// Button B: X+r, Y+s
	// Prize: X=k, Y=l
	//
	// Then
	// no + mr = k
	// np + ms = l
	//
	// From that
	// n = (sk - lr)/(so - rp)
	// m = (l - np)/s
}

type buttomEquasions struct {
	o, p, r, s, k, l int
}

const (
	priceA = 3
	priceB = 1
)

func (d Day13Task1) CalculateAnswer(input string) (string, error) {
	raw := strings.Split(input, "\n")
	var eqs []buttomEquasions
	for i := 0; i < len(raw); i += 4 {
		o, p := d.parseButton(raw[i])
		r, s := d.parseButton(raw[i+1])
		k, l := d.parsePrize(raw[i+2])
		eqs = append(eqs, buttomEquasions{o: o, p: p, r: r, s: s, k: k, l: l})
	}
	totalTokens := 0
	for _, eq := range eqs {
		if n, err := eq.getN(); err != nil {
			log.Println(eq, n, err)
		} else {
			totalTokens += priceA*n + priceB*eq.getMByN(n)
		}
	}
	return strconv.Itoa(totalTokens), nil
}

func (e buttomEquasions) getN() (int, error) {
	divisor := e.s*e.o - e.r*e.p
	dividend := e.s*e.k - e.l*e.r
	if divisor == 0 {
		return -1, fmt.Errorf("Division by 0, there might be multiple solutions")
	}
	n := dividend / divisor
	if n*divisor != dividend {
		return -2, fmt.Errorf("n is a fraction, no solution between whole numbers")
	}
	return n, nil
}

func (e buttomEquasions) getMByN(n int) int {
	return (e.l - n*e.p) / e.s
}

func (d Day13Task1) parseButton(raw string) (int, int) {
	nums := strings.Split(raw[10:], ", ")
	x, _ := strconv.Atoi(nums[0][2:])
	y, _ := strconv.Atoi(nums[1][2:])
	return x, y
}

func (d Day13Task1) parsePrize(raw string) (int, int) {
	nums := strings.Split(raw[7:], ", ")
	x, _ := strconv.Atoi(nums[0][2:])
	y, _ := strconv.Atoi(nums[1][2:])
	return x, y
}
