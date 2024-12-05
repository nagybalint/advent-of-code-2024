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
type pageSet map[page]struct{}
type rules map[page]pageSet
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
	incorrectUpdatePages := make(pageSet)
	incorrectUpdatePages.addAll(ud)
	var correctedUpdate update
	for len(filteredRules) > 0 {
		smallestItem := incorrectUpdatePages.filter(func(p page) bool {
			return !slices.Contains(filteredRules.allAfterPages().pages(), p)
		})[0]
		correctedUpdate = append(correctedUpdate, smallestItem)
		delete(incorrectUpdatePages, smallestItem)
		delete(filteredRules, smallestItem)
	}
	correctedUpdate = append(correctedUpdate, incorrectUpdatePages.pages()[0])
	return correctedUpdate
}

func (rs rules) filterFor(ud update) rules {
	filteredRules := make(rules)
	for before, afters := range rs {
		if !slices.Contains(ud, before) {
			continue
		}
		filteredAfters := utils.Filter(afters.pages(), func(p page) bool { return slices.Contains(ud, p) })
		if len(filteredAfters) == 0 {
			continue
		}
		filteredRules[before] = make(pageSet)
		for _, after := range filteredAfters {
			filteredRules[before].add(after)
		}
	}
	return filteredRules
}

func (rs rules) allAfterPages() pageSet {
	afterPages := make(pageSet)
	for _, after := range rs {
		afterPages.addAllFrom(after)
	}
	return afterPages
}

func (ud update) isCorrectFor(rs rules) bool {
	pagesSeen := make(pageSet)
	for _, newPage := range ud {
		pagesThatShouldComeAfter := rs[newPage]
		for seen := range pagesSeen {
			if pagesThatShouldComeAfter.contains(seen) {
				return false
			}
		}
		pagesSeen.add(newPage)
	}
	return true
}

func (ps pageSet) pages() []page {
	keys := make([]page, len(ps))
	i := 0
	for k := range ps {
		keys[i] = k
		i++
	}
	return keys
}

func (ps pageSet) add(p page) {
	ps[p] = struct{}{}
	return
}

func (ps pageSet) addAllFrom(other pageSet) {
	for p, _ := range other {
		ps.add(p)
	}
}

func (ps pageSet) addAll(others []page) {
	for _, p := range others {
		ps.add(p)
	}
}

func (ps pageSet) filter(test func(p page) bool) []page {
	return utils.Filter(ps.pages(), test)
}

func (ps pageSet) contains(p page) bool {
	_, ok := ps[p]
	return ok
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
			rs[before] = make(pageSet)
		}
		rs[before].add(after)
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
