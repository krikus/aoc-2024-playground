package z5

import (
	"fmt"
	"justatest/utils"
	"slices"
)

type rule struct {
	Value int
	Nexts []int
	Prevs []int
}

func solve1(rules *utils.SortedList[rule], prints [][]int) int {
	total := 0
	for _, p := range prints {
		if isValid(p, rules) {
			total += p[len(p)/2]
		}
	}
	return total
}

func insertIndex(v int, r *utils.SortedList[rule], print []int) int {
  for i := 0; i < len(print); i++ {
    if isValid(append(print[:i], append([]int{v}, print[i:]...)...), r) {
      return i
    }
  }
  return -1
}

func sortByRules(rules *utils.SortedList[rule], print []int) []int {
  sliceCopy := slices.Clone(print)
  slices.SortFunc(sliceCopy, func(a, b int) int {
    ruleA, has := rules.Search(rule{a, nil, nil})
    if !has {
      return 0
    }

    if slices.Contains(rules.Get(ruleA).Nexts, b) {
      return -1
    }
    if slices.Contains(rules.Get(ruleA).Prevs, b) {
      return 1
    }

    return 0
  });
  return sliceCopy
}

func solve2(rules *utils.SortedList[rule], prints [][]int) int {
  total :=0
  for _, p := range prints {
    if !isValid(p, rules) {
      p2 := sortByRules(rules, p)
      total += p2[len(p2)/2]
    }
  }
  return total
}

func buildRules(rules [][]int) utils.SortedList[rule] {
	list := utils.SortedList[rule]{
		Compare: func(a, b rule) int {
			return a.Value - b.Value
		},
	}

	for _, r := range rules {
		a, istherea := list.Search(rule{r[0], nil, nil})
		if istherea {
			list.Get(a).Nexts = append(list.Get(a).Nexts, r[1])
		} else {
			list.Add(rule{r[0], []int{r[1]}, nil})
		}

		b, isthereb := list.Search(rule{r[1], nil, nil})
		if isthereb {
			list.Get(b).Prevs = append(list.Get(b).Prevs, r[0])
		} else {
			list.Add(rule{r[1], nil, []int{r[0]}})
		}
	}

	return list
}

func isValid(prints []int, rules *utils.SortedList[rule]) bool {
	for i, p := range prints {
		ri, found := rules.Search(rule{Value: p})
		if !found {
			continue
		}
		r := rules.Get(ri)
		for j := i + 1; j < len(prints); j++ {
			if slices.Contains(r.Prevs, prints[j]) {
				return false
			}
		}
	}
	return true
}

func Solve() {
	// filepath := "./z5/input-test.txt"
	filepath := "./z5/input.txt"
	lines, _ := utils.ReadLines(filepath)

	rules := [][]int{}

	i := 0
	for i = 0; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		rules = append(rules, utils.ParseNums(lines[i], "|"))
	}

	prints := [][]int{}

	for i = i + 1; i < len(lines); i++ {
		prints = append(prints, utils.ParseNums(lines[i], ","))
	}
	rulesList := buildRules(rules)

	solution := solve1(&rulesList, prints)
	fmt.Printf("SOLUTION: %d\n", solution)
  solution2 := solve2(&rulesList, prints)
  fmt.Printf("SOLUTION2: %d\n", solution2)
}
