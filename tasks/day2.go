package tasks

import (
	"log"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day2Task1 struct{}
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
		if isSafe(r) {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports), nil
}

func isSafe(report Report) bool {
	var diffs []int
	for _, w := range utils.SlidingWindow[Level](report, 2) {
		if len(w) < 2 {
			return true
		}
		diffs = append(diffs, int(w[1])-int(w[0]))
	}
	shouldIncrease := diffs[0] > 0
	for _, d := range diffs {
		if shouldIncrease {
			if d < 0 {
				return false
			}
		} else {
			if d > 0 {
				return false
			}
		}
		if distance := utils.AbsInt(d); distance < 1 {
			return false
		} else if distance > 3 {
			return false
		}
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
