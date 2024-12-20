package z4

import (
	"fmt"
	"justatest/utils"
)

func checkXmass(x, y int, lines []string) int {
	letters := []byte{'M', 'A', 'S'}
	found := 0

	for xes := -1; xes <= 1; xes++ {
		for yes := -1; yes <= 1; yes++ {
			if xes == 0 && yes == 0 {
				continue
			}
			newx := x
			newy := y
			i := 0
			for i = 0; i < len(letters); i++ {
				newx += xes
				newy += yes
				if newx < 0 || newy < 0 || newx >= len(lines[0]) || newy >= len(lines) {
					break
				}
				if lines[newy][newx] != letters[i] {
					break
				}
			}
			if i == len(letters) {
				found++
			}
		}
	}
	return found
}

func isLetter(c byte) bool {
	return c == 'M' || c == 'S'
}

func checkMas(x, y int, lines []string) int {
	if x == 0 || y == 0 || x >= len(lines[0])-1 || y >= len(lines)-1 {
		return 0
	}

	pairs := [][]int{
		{y + 1, x + 1, y - 1, x - 1},
		{y - 1, x + 1, y + 1, x - 1},
	}

	for _, pair := range pairs {
		if isLetter(lines[pair[0]][pair[1]]) && isLetter(lines[pair[2]][pair[3]]) && lines[pair[0]][pair[1]] != lines[pair[2]][pair[3]] {
      continue
		}
    return 0
	}
  return 1
}

func Solve() {
	// filepath := "./z4/input-test.txt"
	filepath := "./z4/input.txt"
	lines, _ := utils.ReadLines(filepath)

	queueX := utils.SortedList[utils.Point]{
		Compare: func(a, b utils.Point) int {
			return a.Y - b.Y
		},
	}

	queueA := utils.SortedList[utils.Point]{
		Compare: func(a, b utils.Point) int {
			return a.X - b.X
		},
	}

	for idy, line := range lines {
		for idx, char := range line {
			if char == 'X' {
				queueX.Add(utils.Point{X: idx, Y: idy})
			}
			if char == 'A' {
				queueA.Add(utils.Point{X: idx, Y: idy})
			}

		}
	}

	total1 := 0
	for queueX.IsEmpty() == false {
		point := queueX.Pop()
		total1 += checkXmass(point.X, point.Y, lines)
	}

	fmt.Printf("SOLUTION: %d\n", total1)

  total2 := 0
  for queueA.IsEmpty() == false {
    point := queueA.Pop()
    total2 += checkMas(point.X, point.Y, lines)
  }
  fmt.Printf("SOLUTION: %d\n", total2)
}
