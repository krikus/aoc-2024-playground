package z6

import (
	"fmt"
	"justatest/utils"
)

func solve1(map2d *utils.Map2d[byte]) int {
  positions := []byte{'^', '>', 'v', '<'}
  dirs := []utils.Point{
    utils.NewPoint(0, -1),
    utils.NewPoint(1, 0),
    utils.NewPoint(0, 1),
    utils.NewPoint(-1, 0),
  }

  dir := 0
  total := 1
  startPoint := map2d.FindPosition(positions[dir])
  for startPoint == nil && dir < 4 {
    dir++;
    startPoint = map2d.FindPosition(positions[dir])
  }
  
  fmt.Printf("startPoint: %v [%v/%d]\n", startPoint, positions[dir], dir)

  for {
    if startPoint == nil {
      return total;
    }

    nextPoint := startPoint.Clone().Add(&dirs[dir])
    nextByte, exists := map2d.Get(nextPoint)
    if (!exists) {
      break
    }
    if *nextByte == '#' {
      dir = (dir + 1) % 4
      continue
    } else if *nextByte == '.' {
      total++
    } 
    map2d.Set('X', startPoint)
    startPoint = nextPoint
  }

  return total
}

func Solve() {
	// filepath := "./z6/input-test.txt"
	filepath := "./z6/input.txt"
	lines, _ := utils.ReadLines(filepath)

  all_lines := make([][]byte, len(lines))

  for i, line := range lines {
    all_lines[i] = []byte(line)
  }

  map2d := utils.CreateMap2d(all_lines)
  fmt.Printf("solution #1: %d\n", solve1(map2d))
}
