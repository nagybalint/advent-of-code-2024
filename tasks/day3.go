package tasks

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day3Task1 struct{}

func (Day3Task1) CalculateAnswer() (string, error) {
	input, err := utils.ReadFileFromRelative("resources/day3.txt")
	if err != nil {
		log.Println("Cannot open input file", err)
		return "", nil
	}
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	sum := 0
	for _, m := range matches {
		log.Println(m)
		operands := strings.Split(strings.Trim(m, "mul()"), ",")
		if len(operands) != 2 {
			return "", fmt.Errorf("Expected 2 operands, instead got ", len(operands))
		}
		a, err := strconv.Atoi(operands[0])
		if err != nil {
			log.Println("Cannot convert ", operands[0], " to integer", err)
			return "", err
		}
		b, err := strconv.Atoi(operands[1])
		if err != nil {
			log.Println("Cannot convert ", operands[1], " to integer", err)
			return "", err
		}
		sum += a * b
	}
	return strconv.Itoa(sum), nil
}

type mulExpression string
