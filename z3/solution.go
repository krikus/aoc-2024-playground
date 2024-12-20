package z3

import (
	"fmt"
	"justatest/utils"
	"regexp"
	"strconv"
	"strings"
)

// enum declaration
type job int

const (
	mul job = iota
	do
	dont
)

type instruction struct {
	instr job
	pair  [2]int
}

func parseRegex(line string) []instruction {
	reg := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)

	found := reg.FindAllStringSubmatch(line, -1)
	instructions := make([]instruction, 0)
	for _, match := range found {
		if strings.Contains(match[0], "mul") {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			pairs := [2]int{a, b}
			instructions = append(instructions, instruction{instr: mul, pair: pairs})
		}
		if strings.Contains(match[0], "do") {
			instructions = append(instructions, instruction{instr: do})
		}
		if strings.Contains(match[0], "don't") {
			instructions = append(instructions, instruction{instr: dont})
		}
	}

	return instructions
}

func solve1(all []instruction, forceDo bool) int {
	total := 0
  switched := true
	for _, instr := range all {
		if instr.instr == mul && (switched || forceDo) {
			total += instr.pair[0] * instr.pair[1]
		} else if instr.instr == do {
      switched = true
		} else if instr.instr == dont {
      switched = false
		}
	}
	return total
}

func Solve() {
	// filepath := "./z3/input-test.txt"
	filepath := "./z3/input.txt"

	lines, _ := utils.ReadLines(filepath)

	pairs := make([]instruction, 0)
	for _, line := range lines {
		pairs = append(pairs, parseRegex(line)...)
	}
	solution := solve1(pairs, true)
	fmt.Printf("SOLUTION: %d\n", solution)
  solution = solve1(pairs, false)
  fmt.Printf("SOLUTION: %d\n", solution)
}
