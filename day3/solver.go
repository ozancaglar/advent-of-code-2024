package day3

import (
	"log"
	"regexp"
	"strings"

	"github.com/ozancaglar/advent-of-code-2024/util"
)

func Solve(filename string) {
	parts(filename)
}

func parts(filename string) {
	scanner := util.StreamLines(filename)
	re := regexp.MustCompile("mul\\(([0-9]{1,3},[0-9]{1,3})\\)")
	p2re := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don't\\(\\)")
	p1Total := 0
	p2Total := 0

	do := true

	for scanner.Scan() {
		for _, str := range re.FindAllString(scanner.Text(), -1) {
			p1Total += extractAdditionFromMult(str)
		}

		for _, str := range p2re.FindAllString(scanner.Text(), -1) {
			if str == "do()" {
				do = true
				continue
			}

			if str == "don't()" {
				do = false
				continue
			}

			if !do {
				continue
			}

			p2Total += extractAdditionFromMult(str)
		}
	}

	log.Printf("Day 3 Part 1 total: %d", p1Total)
	log.Printf("Day 3 Part 2 total: %d", p2Total)

}

func extractAdditionFromMult(str string) int {
	newString := strings.TrimPrefix(str, "mul(")
	newString = strings.TrimSuffix(newString, ")")
	left, right := strings.Split(newString, ",")[0], strings.Split(newString, ",")[1]

	leftInt, rightInt := util.MustParseInt(left), util.MustParseInt(right)
	return leftInt * rightInt
}
