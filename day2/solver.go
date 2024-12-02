package day2

import (
	"log"
	"slices"
	"strings"

	"github.com/ozancaglar/advent-of-code-2024/util"
)

func Solve(filename string) {
	parts(filename)
}

func parts(filename string) {
	lines := util.StreamLines(filename)

	partOneTotal := 0
	partTwoTotal := 0

lineLoop:
	for lines.Scan() {
		numbers := make([]int, 0)
		for _, number := range strings.Split(lines.Text(), " ") {
			numbers = append(numbers, util.MustParseInt(number))
		}

		if !isNotSafe(numbers) {
			partOneTotal++
		}

		for i := 0; i < len(numbers); i++ {
			newSlice := slices.Clone(numbers)
			newSlice = append(newSlice[:i], newSlice[i+1:]...)
			if !isNotSafe(newSlice) {
				partTwoTotal++
				continue lineLoop
			}
		}
	}
	log.Printf("Part 1: %d", partOneTotal)
	log.Printf("Part 2: %d", partTwoTotal)
}

func isNotSafe(numbers []int) bool {
	alwaysDecreasing := false
	alwaysIncreasing := false
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			continue
		}
		diff := numbers[i] - numbers[i-1]
		if diff == 0 {
			return true
		}

		if diff > 3 || diff < -3 {
			return true
		}

		if diff < 0 {
			alwaysDecreasing = true
		}

		if diff > 0 {
			alwaysIncreasing = true
		}

		if alwaysDecreasing && alwaysIncreasing {
			return true
		}

	}
	return false
}
