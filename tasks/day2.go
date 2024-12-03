package tasks

import (
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day2Task1 struct{}
type Day2Task2 struct{}
type Level int
type Report []Level

func (Day2Task1) CalculateAnswer() (string, error) {
	input, err := utils.ReadFileFromRelative("resources/day2.txt")
	if err != nil {
		log.Println("Error reading input")
		return "", err
	}

	reports, err := getReports(input)
	if err != nil {
		log.Println("Cannot parse all reports", err)
		return "", err
	}

	safeReports := 0
	for _, r := range reports {
		if isSafe(r, true) || isSafe(r, false) {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports), nil
}

func (Day2Task2) CalculateAnswer() (string, error) {
	input, err := utils.ReadFileFromRelative("resources/day2.txt")
	if err != nil {
		log.Println("Error reading input")
		return "", err
	}

	reports, err := getReports(input)
	if err != nil {
		log.Println("Cannot parse all reports", err)
		return "", err
	}

	safeReports := 0
	for _, r := range reports {
		if len(r) < 2 {
			safeReports++
		} else if isSafeWithToleration(r, true, true) ||
			isSafeWithToleration(r, false, true) ||
			isSafeWithToleration(r[1:], true, false) ||
			isSafeWithToleration(r[1:], false, false) {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports), nil
}

func isSafeWithToleration(report Report, isIncreasing bool, isFailureTolerated bool) bool {
	if len(report) == 1 {
		return true
	}
	if len(report) == 2 {
		if isFailureTolerated {
			return true
		} else {
			return isStepValid(report[0], report[1], isIncreasing)
		}
	}

	// At this point the length of the report is at least 3
	if isStepValid(report[0], report[1], isIncreasing) {
		return isSafeWithToleration(report[1:], isIncreasing, isFailureTolerated)
	}
	if !isFailureTolerated {
		return false
	}
	if isStepValid(report[0], report[2], isIncreasing) {
		return isSafeWithToleration(report[2:], isIncreasing, false)
	}
	return false
}

// Tail recursive call -- equivalent to isSafeWithToleration(report, isIncreasing, false),
// kept here because this was originally used to solve task 1
func isSafe(report Report, isIncreasing bool) bool {
	if len(report) < 2 {
		return true
	}
	if !isStepValid(report[0], report[1], isIncreasing) {
		return false
	}
	return isSafe(report[1:], isIncreasing)
}

func isStepValid(a, b Level, isIncreasing bool) bool {
	diff := b - a
	if isIncreasing {
		if diff < 0 {
			return false
		}
	} else {
		if diff > 0 {
			return false
		}
	}
	if distance := utils.AbsInt(int(diff)); distance < 1 {
		return false
	} else if distance > 3 {
		return false
	}
	return true
}

func getReports(input string) ([]Report, error) {
	lines := utils.Filter(strings.Split(input, "\n"), func(s string) bool {
		return s != ""
	})
	var reports []Report
	for _, line := range lines {
		var levels []Level
		rawLevels := strings.Split(strings.TrimRight(line, "\n"), " ")
		for _, l := range rawLevels {
			if level, err := strconv.Atoi(l); err != nil {
				log.Println("Cannot parse level", err)
				return nil, err
			} else {
				levels = append(levels, Level(level))
			}
		}
		reports = append(reports, levels)
	}
	return reports, nil
}
