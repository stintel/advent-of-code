package main

import "bufio"
import "log"
import "os"
import "regexp"
import "strconv"
import "strings"

type Rule struct {
	Amount int
	Color string
}

type Bag struct {
	Color string
	Rules []Rule
}

func parseRule(rs string) (bag Bag) {
	var rule Rule

	t := strings.SplitAfter(rs, "contain")

	bag.Color = strings.Replace(t[0], " bags contain", "", 1)

	t = strings.SplitAfter(t[1], ",")
	for _, r := range t {
		r = strings.TrimLeft(r, " ")
		if r == "no other bags." {
			return
		}
		reg := regexp.MustCompile(` bag.*[\.|\,]$`)
		r = reg.ReplaceAllString(r, "")
		ra := strings.SplitN(r, " ", 2)
		rule.Amount, _ = strconv.Atoi(ra[0])
		rule.Color = ra[1]

		bag.Rules = append(bag.Rules, rule)
	}

	return
}

func readFile(path string) (data []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func mergeRules(Bags []Bag) (newbags []Bag) {
	var newrules []Rule
	var newbag Bag

	for _, b := range Bags {
		newrules = nil
		//log.Printf("Bag color: %s", b.Color)
		//log.Printf("Old rules: %s", b.Rules)
		for _, r := range b.Rules {
			if r.Color == "shiny gold" {
				newrules = append(newrules, r)
			}
			for _, bt := range Bags {
				if r.Color == bt.Color {
					for _, rt := range bt.Rules {
						rt.Amount = r.Amount * rt.Amount
						newrules = append(newrules, rt)
					}
				}
			}
		}
		newbag.Color = b.Color
		newbag.Rules = newrules

		if len(newbag.Rules) > 0 {
			newbags = append(newbags, newbag)
		}
	}

	return
}

func countBags(bags []Bag, color string) (count int) {

	for _, b := range bags {
		if b.Color == color {
			for _, r := range b.Rules {
				count += r.Amount
				count += r.Amount * countBags(bags, r.Color)
			}
		}
	}

	return
}

func main() {
	var Bags []Bag
	var numrules int = 0

	testdata := readFile("input")

	for _, line := range testdata {
		Bags = append(Bags, parseRule(line))
	}

	newbags := mergeRules(Bags)
	numrules = len(newbags)

	for {
		newbags = mergeRules(newbags)

		if numrules != len(newbags) {
			numrules = len(newbags)
		} else {
			break
		}
	}

	log.Printf("Number of new rules: %d", len(newbags))
	log.Printf("Number of bags contained in a single shiny gold bag: %d", countBags(Bags, "shiny gold"))
}
