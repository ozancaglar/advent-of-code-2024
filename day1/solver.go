package day1

import (
	"github.com/ozancaglar/advent-of-code-2024/util"
	"log"
	"regexp"
	"slices"
	"strconv"
)

func Solve(filename string) {
	parts(filename)
}

func parts(filename string) {
	scanner := util.StreamLines(filename)

	re := regexp.MustCompile("[0-9]+")
	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)
	appearancesInRightList := make(map[int]int)
	total := 0
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		left, _ := strconv.Atoi(matches[0])
		right, _ := strconv.Atoi(matches[1])
		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
		appearancesInRightList[right]++
	}

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)
	if len(leftSlice) != len(rightSlice) {
		log.Fatal("Left and right slices are not the same length")
	}

	for i := 0; i < len(leftSlice); i++ {
		diff := leftSlice[i] - rightSlice[i]
		if diff < 0 {
			diff = diff * -1
		}
		total += diff
	}

	log.Printf("Day one, part one answer: %v", total)

	totalPartTwo := 0
	for i := 0; i < len(leftSlice); i++ {
		appearancesInRightListForThisInt := appearancesInRightList[leftSlice[i]]
		if appearancesInRightListForThisInt > 0 {
			totalPartTwo += leftSlice[i] * appearancesInRightListForThisInt
		}
	}
	log.Printf("Day one, part two answer: %v", totalPartTwo)

}
