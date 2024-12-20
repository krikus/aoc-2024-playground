package z2

import (
	"fmt"
	"justatest/utils"
	"os"
	"strconv"
	"strings"
)

type orderingChecker = func(int, int, int) bool

func parseNums(line string) []int {
	nums := []int{}
	for _, str := range strings.Split(line, " ") {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func solveTask(all [][]int, maxdiff int, canRetry bool) int {
	safeReps := 0
	for _, nums := range all {
		dir := nums[len(nums)-1] - nums[0]

		var isOrdered orderingChecker

		if dir > 0 {
			isOrdered = isIncreasing
		} else if dir < 0 {
			isOrdered = isDecreasing
		} else {
			continue
		}

		wasGood := checkOrder(nums, maxdiff, isOrdered, canRetry)
		if wasGood {
			safeReps++
		}
	}
	return safeReps
}

func copyArrayWithoutIndex(nums []int, idx int) []int {
	newArray := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != idx {
			newArray = append(newArray, nums[i])
		}
	}
	return newArray
}

func checkOrder(nums []int, maxdiff int, order orderingChecker, retry bool) bool {
	for idx := 0; idx < len(nums)-1; idx++ {
		if !order(nums[idx], nums[idx+1], maxdiff) {
			if retry {
				retryOne := copyArrayWithoutIndex(nums, idx)
				if checkOrder(retryOne, maxdiff, order, false) {
					return true
				}
				retryOne = copyArrayWithoutIndex(nums, idx+1)
				if checkOrder(retryOne, maxdiff, order, false) {
					return true
				}

			}
			return false
		}
	}

	return true
}

func isIncreasing(a, b, maxdiff int) bool {
	diff := b - a
	return diff > 0 && diff <= maxdiff
}

func isDecreasing(a, b, maxdiff int) bool {
	diff := a - b
	return diff > 0 && diff <= maxdiff
}

func Solve() {
	// filepath := "./z2/input-2a-test.txt"
	filepath := "./z2/input-2a.txt"

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var all [][]int = make([][]int, 0)

	allNums, _ := utils.ReadFile(filepath, parseNums)
	for _, nums := range allNums {
		if len(nums) > 1 {
			all = append(all, nums)
		}
	}

	defer file.Close()

	solution := solveTask(all, 3, false)
	fmt.Printf("SOLUTION1: %d\n", solution)
  solution2 := solveTask(all, 3, true)
  fmt.Printf("SOLUTION2: %d\n", solution2)
}
