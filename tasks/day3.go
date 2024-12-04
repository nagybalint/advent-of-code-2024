package tasks

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Day3Task1 struct{}

func (Day3Task1) CalculateAnswer(input string) (string, error) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	sum := 0
	for _, m := range matches {
		result, err := calculate(mulExpression(m))
		if err != nil {
			log.Println("Cannot calculate result of mul expression", err)
			return "", err
		}
		sum += result
	}
	return strconv.Itoa(sum), nil
}

type Day3Task2 struct{}
type match struct {
	index   int
	content string
}

func (Day3Task2) CalculateAnswer(input string) (string, error) {
	mulExpressions, err := findAllWithIndex(`mul\(\d+,\d+\)`, input)
	if err != nil {
		log.Println("Error while parsing mul expressions")
		return "", nil
	}

	dos, err := findAllWithIndex(`do\(\)`, input)
	if err != nil {
		log.Println("Error while parsing dos")
		return "", nil
	}

	donts, err := findAllWithIndex(`don't\(\)`, input)
	if err != nil {
		log.Println("Error while parsing dos")
		return "", nil
	}

	allExpressions := slices.Concat[[]match](mulExpressions, dos, donts)
	sort.Slice(allExpressions, func(i, j int) bool {
		return allExpressions[i].index < allExpressions[j].index
	})

	sum := 0
	enabled := true
	for _, expr := range allExpressions {
		switch expr.content {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				res, err := calculate(mulExpression(expr.content))
				if err != nil {
					log.Println("Cannot calculate result of mul expression", err)
					return "", err
				}
				sum += res
			} else {
				continue
			}
		}
	}

	return strconv.Itoa(sum), nil
}

type mulExpression string

func findAllWithIndex(regex string, s string) ([]match, error) {
	reMul := regexp.MustCompile(regex)
	strs := reMul.FindAllString(s, -1)
	indices := reMul.FindAllStringIndex(s, -1)
	if len(strs) != len(indices) {
		return nil, fmt.Errorf("Expected as many matches as match indices")
	}
	var matches []match
	for i := range strs {
		matches = append(matches, match{
			index:   indices[i][0],
			content: strs[i],
		})
	}
	return matches, nil
}
func calculate(expression mulExpression) (int, error) {
	operands := strings.Split(strings.Trim(string(expression), "mul()"), ",")
	if len(operands) != 2 {
		return -1, fmt.Errorf("Expected 2 operands, instead got ", len(operands))
	}
	a, err := strconv.Atoi(operands[0])
	if err != nil {
		log.Println("Cannot convert ", operands[0], " to integer", err)
		return -1, err
	}
	b, err := strconv.Atoi(operands[1])
	if err != nil {
		log.Println("Cannot convert ", operands[1], " to integer", err)
		return -1, err
	}
	return a * b, nil
}
