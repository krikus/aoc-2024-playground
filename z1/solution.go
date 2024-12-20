package z1

import (
	"fmt"
	"justatest/utils"
	"math"
	"os"
)

type SortedList = utils.SortedList[int]

func parseNums(line string) []int {
  return utils.ParseNums(line, " ")
}

func solve1(list1 *SortedList, list2 *SortedList) int {
	total := 0
	for {
		num1 := list1.Pop()
		num2 := list2.Pop()
		total += int(math.Abs(float64(num1 - num2)))
		if list1.IsEmpty() {
			break
		}
	}
	return total
}

func solve2(list1 *SortedList, list2 *SortedList) int {
	total := 0
	for !list1.IsEmpty() {
		times := 0
		item := list1.Pop()
		for !list2.IsEmpty() && item >= list2.Peek() {
			l2Item := list2.Pop()
			if item == l2Item {
				times++
			}
		}
		lastScore := times * item
		if lastScore > 0 {
			total += lastScore
			for !list1.IsEmpty() && item == list1.Peek() {
				list1.Pop()
				total += lastScore
			}
		}
	}
	return total
}

func Solve() {
	//filepath := "./z1/input-1a-test.txt"
	filepath := "./z1/input-1a.txt"

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	less := func(a, b int) int {
		return a - b
	}

	firstList := &SortedList{
		Compare: less}
	secondList := &SortedList{
		Compare: less}

	allNums, _ := utils.ReadFile(filepath, parseNums)
	for _, nums := range allNums {
		if len(nums) == 2 {
			firstList.Add(nums[0])
			secondList.Add(nums[1])
			continue
		}
	}

	defer file.Close()

	solution := solve1(firstList.Clone(), secondList.Clone())
	solution2 := solve2(firstList, secondList)
	fmt.Printf("SOLUTION1: %d\n", solution)
	fmt.Printf("SOLUTION2: %d\n", solution2)
}
