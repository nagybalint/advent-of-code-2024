package tasks

import (
	"fmt"
	"log"
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
