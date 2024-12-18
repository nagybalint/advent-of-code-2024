package tasks

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day1Task1 struct{}

func (Day1Task1) CalculateAnswer(input string) (string, error) {
	locationsA, locationsB, err := getLocationIds(input)
	if err != nil {
		log.Println("Cannot parse location ids")
		return "", err
	}

	sort.Ints(locationsA)
	sort.Ints(locationsB)

	distance := getTotalDistance(locationsA, locationsB)

	return strconv.Itoa(distance), nil
}

type Day1Task2 struct{}

func (Day1Task2) CalculateAnswer(input string) (string, error) {
	locationsA, locationsB, err := getLocationIds(input)
	if err != nil {
		log.Println("Cannot parse location ids")
		return "", err
	}
	rightOccurrences := getOccurences(locationsB)
	similarity := getSimilarity(locationsA, rightOccurrences)
	return strconv.Itoa(similarity), nil
}

func getSimilarity(leftLocations []int, rightOccurrences map[int]int) (similarity int) {
	for _, leftLocation := range leftLocations {
		similarity += leftLocation * rightOccurrences[leftLocation]
	}
	return
}

func getOccurences(locations []int) map[int]int {
	occurrences := make(map[int]int)
	for _, locationId := range locations {
		occurrences[locationId]++
	}
	return occurrences
}

func parseAndAppendLocationId(list []int, id string) ([]int, error) {
	locationId1, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Cannot convert locationId to integer", err)
		return nil, err
	}
	list = append(list, locationId1)
	return list, nil
}

func getLocationIds(input string) ([]int, []int, error) {
	lines := utils.Filter(strings.Split(input, "\n"), func(s string) bool {
		return s != ""
	})

	var locationsA, locationsB []int

	re := regexp.MustCompile(`(\d*) *(\d*) *`)
	for _, line := range lines {
		captureGroups := re.FindStringSubmatch(line)
		var err error
		if locationsA, err = parseAndAppendLocationId(locationsA, captureGroups[1]); err != nil {
			return nil, nil, err
		}
		if locationsB, err = parseAndAppendLocationId(locationsB, captureGroups[2]); err != nil {
			return nil, nil, err
		}
	}

	return locationsA, locationsB, nil
}

func getTotalDistance(locationsA, locationsB []int) int {
	if len(locationsA) != len(locationsB) {
		log.Fatal("Lists are not of the same length")
	}

	var distance int = 0

	for i := range locationsA {
		log.Println(utils.AbsInt(locationsA[i] - locationsB[i]))
		distance += utils.AbsInt(locationsA[i] - locationsB[i])
	}

	return distance
}
