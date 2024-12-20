package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile[T any](path string, mapper func(string) []T) ([][]T, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var lines [][]T = make([][]T, 0)
  for scanner.Scan() {
    lines = append(lines, mapper(scanner.Text()))
  }
  
  return lines, nil
}

func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var lines []string = make([]string, 0)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  
  return lines, nil
}


func ParseNums(line string, splitBy string) []int {
	nums := []int{}
	for _, str := range strings.Split(line, splitBy) {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}
