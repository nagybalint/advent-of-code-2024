package tasks

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/nagybalint/advent-of-code-2024/utils"
)

type Day5Task1 struct{}
type page int
type rules map[page]utils.Set[page]
type update []page

func (Day5Task1) CalculateAnswer(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		return "", fmt.Errorf("Expected the puzzle input to contain 2 parts separated by an empty line")
	}

	rawRules := utils.Filter(strings.Split(parts[0], "\n"), utils.IsNonEmptyString)
	rawUpdates := utils.Filter(strings.Split(parts[1], "\n"), utils.IsNonEmptyString)

	rules, err := parseRules(rawRules)
	if err != nil {
		log.Println("Error parsing raw rules", err)
		return "", err
	}
	updates, err := parseUpdates(rawUpdates)
	if err != nil {
		log.Println("Cannot parse raw updates", err)
		return "", err
	}

	log.Println(rules)
	log.Println(updates)

	var sum int
	for i, correctUpdate := range utils.Filter(updates, func(u update) bool { return u.isCorrectFor(rules) }) {
		log.Println("Correct update at ", i, " with middle element ", correctUpdate.getMiddlePage())
		sum += int(correctUpdate.getMiddlePage())
	}

	return strconv.Itoa(sum), nil
}

type Day5Task2 struct{}

// As there are n*(n-1)/2 sorting rules in the task input
// (the relation of any 2 numbers appearing in the updates is defined),
// a simple bubble sort with a custom comparator would work as well
// Building up a graph based on the ordering rules would also work
func (Day5Task2) CalculateAnswer(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		return "", fmt.Errorf("Expected the puzzle input to contain 2 parts separated by an empty line")
	}

	rawRules := utils.Filter(strings.Split(parts[0], "\n"), utils.IsNonEmptyString)
	rawUpdates := utils.Filter(strings.Split(parts[1], "\n"), utils.IsNonEmptyString)

	rules, err := parseRules(rawRules)
	if err != nil {
		log.Println("Error parsing raw rules", err)
		return "", err
	}
	updates, err := parseUpdates(rawUpdates)
	if err != nil {
		log.Println("Cannot parse raw updates", err)
		return "", err
	}

	log.Println(rules)
	log.Println(updates)

	incorrectUpdates := utils.Filter(updates, func(u update) bool { return !u.isCorrectFor(rules) })
	var correctedUpdates []update
	for _, incorrectUpdate := range incorrectUpdates {
		filteredRules := rules.filterFor(incorrectUpdate)
		log.Println("update ", incorrectUpdate, " filtered rules ", filteredRules)
		correctedUpdate := incorrectUpdate.correct(filteredRules)
		correctedUpdates = append(correctedUpdates, correctedUpdate)
	}

	var sum int
	for _, ud := range correctedUpdates {
		sum += int(ud.getMiddlePage())
	}

	return strconv.Itoa(sum), nil
}

func (ud update) correct(filteredRules rules) update {
	incorrectUpdatePages := make(utils.Set[page])
	incorrectUpdatePages.AddAll(ud)
	var correctedUpdate update
	for len(filteredRules) > 0 {
		smallestItem := filterPageSet(incorrectUpdatePages, func(p page) bool {
			return !slices.Contains(filteredRules.allAfterPages().ToSlice(), p)
		})[0]
		correctedUpdate = append(correctedUpdate, smallestItem)
		delete(incorrectUpdatePages, smallestItem)
		delete(filteredRules, smallestItem)
	}
	correctedUpdate = append(correctedUpdate, incorrectUpdatePages.ToSlice()[0])
	return correctedUpdate
}

func (rs rules) filterFor(ud update) rules {
	filteredRules := make(rules)
	for before, afters := range rs {
		if !slices.Contains(ud, before) {
			continue
		}
		filteredAfters := utils.Filter(afters.ToSlice(), func(p page) bool { return slices.Contains(ud, p) })
		if len(filteredAfters) == 0 {
			continue
		}
		filteredRules[before] = make(utils.Set[page])
		for _, after := range filteredAfters {
			filteredRules[before].Add(after)
		}
	}
	return filteredRules
}

func (rs rules) allAfterPages() utils.Set[page] {
	afterPages := make(utils.Set[page])
	for _, after := range rs {
		afterPages.AddAll(after.ToSlice())
	}
	return afterPages
}

func (ud update) isCorrectFor(rs rules) bool {
	pagesSeen := make(utils.Set[page])
	for _, newPage := range ud {
		pagesThatShouldComeAfter := rs[newPage]
		for seen := range pagesSeen {
			if pagesThatShouldComeAfter.Contains(seen) {
				return false
			}
		}
		pagesSeen.Add(newPage)
	}
	return true
}

func filterPageSet(ps utils.Set[page], test func(p page) bool) []page {
	return utils.Filter(ps.ToSlice(), test)
}

func (ud update) getMiddlePage() page {
	return ud[len(ud)/2]
}

func parseUpdates(rawUpdates []string) ([]update, error) {
	var updates []update
	for _, rawUpdateLine := range rawUpdates {
		var ud update
		rawPages := strings.Split(rawUpdateLine, ",")
		for _, rp := range rawPages {
			_p, err := strconv.Atoi(rp)
			if err != nil {
				log.Println("Cannot parse page", err)
				return nil, err
			}
			ud = append(ud, page(_p))
		}
		updates = append(updates, ud)
	}
	return updates, nil
}

func parseRules(rawRules []string) (rules, error) {
	rs := make(rules)
	for _, r := range rawRules {
		before, after, err := parseRule(r)
		if err != nil {
			log.Println("Error while parsing rule ", r, " ", err)
			return nil, err
		}
		_, ok := rs[before]
		if !ok {
			rs[before] = make(utils.Set[page])
		}
		rs[before].Add(after)
	}
	return rs, nil
}

func parseRule(rawRule string) (before page, after page, err error) {
	ruleParts := strings.Split(rawRule, "|")
	if len(ruleParts) != 2 {
		err = fmt.Errorf("Expected the rule to have 2 parts separated by |")
		return
	}
	var _p int
	_p, err = strconv.Atoi(ruleParts[0])
	if err != nil {
		log.Println("Invalid page value", err)
		return
	}
	before = page(_p)
	_p, err = strconv.Atoi(ruleParts[1])
	if err != nil {
		log.Println("Invalid page value", err)
		return
	}
	after = page(_p)
	return
}
