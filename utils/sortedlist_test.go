package utils_test

import (
  "justatest/utils"
  "testing"
)

func TestSortedList(t *testing.T) {
  list := utils.SortedList[int]{
    Compare: func(a, b int) int {
      return a - b
    },
  }

  list.Add(1)
  list.Add(3)
  list.Add(2)

  nums := []int{}
  for !list.IsEmpty() {
    nums = append(nums, list.Pop())
  }

  if nums[0] != 1 {
    t.Errorf("Expected 1, got %d", nums[0])
  }
  if nums[1] != 2 {
    t.Errorf("Expected 2, got %d", nums[1])
  }
  if nums[2] != 3 {
    t.Errorf("Expected 3, got %d", nums[2])
  }
}

func TestSearch(t *testing.T) {
  list := utils.SortedList[int]{
    Compare: func(a, b int) int {
      return a - b
    },
  }

  list.Add(1)
  list.Add(5)
  list.Add(3)
  list.Add(2)


  index, found := list.Search(5)

  if index != 3 {
    t.Errorf("Expected 3, got %d", index)
  }
  if found != true {
    t.Errorf("Expected true, got %v", found)
  }


  index, found = list.Search(2)

  if index != 1 {
    t.Errorf("Expected 1, got %d", index)
  }
  if found != true {
    t.Errorf("Expected true, got %v", found)
  }


  _, found = list.Search(0)

  if found != false {
    t.Errorf("Expected false, got %v", found)
  }
}
