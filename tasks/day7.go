package tasks

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day7Task1 struct{}
type operation func(int64, int64) int64

type calibration struct {
	result   int64
	operands []int64
}

func (Day7Task1) CalculateAnswer(input string) (string, error) {
	var calibrations []calibration
	for _, line := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		calibration, err := parseCalibration(line)
		if err != nil {
			log.Println("Cannot parse calibration", err)
			return "", nil
		}
		calibrations = append(calibrations, *calibration)
	}

	var totalResult int64 = 0
	for _, c := range calibrations {
		if c.canBeMadeTrue([]operation{add, mul}) {
			log.Println(c)
			totalResult += c.result
		}
	}

	return fmt.Sprintf("%v", totalResult), nil
}

type Day7Task2 struct{}

func (Day7Task2) CalculateAnswer(input string) (string, error) {
	var calibrations []calibration
	for _, line := range utils.Filter(strings.Split(input, "\n"), utils.IsNonEmptyString) {
		calibration, err := parseCalibration(line)
		if err != nil {
			log.Println("Cannot parse calibration", err)
			return "", nil
		}
		calibrations = append(calibrations, *calibration)
	}

	var totalResult int64 = 0
	for _, c := range calibrations {
		if c.canBeMadeTrue([]operation{add, mul, concat}) {
			log.Println(c)
			totalResult += c.result
		}
	}

	return fmt.Sprintf("%v", totalResult), nil
}

func (c calibration) canBeMadeTrue(allowedOps []operation) bool {
	return helper(c.result, c.operands[1:], c.operands[0], allowedOps)
}

func helper(target int64, remainingOperands []int64, alreadyCounted int64, allowedOps []operation) bool {
	if len(remainingOperands) == 0 {
		return alreadyCounted == target
	}
	for _, op := range allowedOps {
		if helper(target, remainingOperands[1:], op(alreadyCounted, remainingOperands[0]), allowedOps) {
			return true
		}
	}
	return false
}

func add(x, y int64) int64 {
	return x + y
}

func mul(x, y int64) int64 {
	return x * y
}

func concat(x, y int64) int64 {
	concatted, _ := strconv.ParseInt(fmt.Sprintf("%v%v", x, y), 10, 64)
	return concatted
}

func parseCalibration(input string) (*calibration, error) {
	re := regexp.MustCompile(`\d+`)
	nums := re.FindAllString(input, -1)
	result, err := strconv.ParseInt(nums[0], 10, 64)
	if err != nil {
		log.Println("Cannot parse calibration result", err)
		return nil, err
	}

	var operands []int64
	for _, num := range nums[1:] {
		op, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Println("Cannot parse operand", err)
			return nil, err
		}
		operands = append(operands, op)
	}

	return &calibration{
		result:   result,
		operands: operands,
	}, nil
}
