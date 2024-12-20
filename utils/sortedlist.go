package utils

import "slices"

type SortedList[T any] struct {
	items []T
	Compare func(T, T) int
}

func (s *SortedList[T]) Add(item T) {
  index, _ := s.Search(item)
  s.items = slices.Insert(s.items, index, item)
}

func (s *SortedList[T]) Get(index int) *T {
  return &s.items[index]
}

// Search returns the index of the item in the list, or index for insertion
func (s *SortedList[T]) Search(searchFor T) (int, bool) {
  index, found := slices.BinarySearchFunc(s.items, searchFor, s.Compare)
  return index, found
}

func (s *SortedList[T]) Peek() T {
	return s.items[0]
}

func (s *SortedList[T]) Pop() T {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *SortedList[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *SortedList[T]) Clone() *SortedList[T] {
	items := append([]T{}, s.items...)
	clone := &SortedList[T]{
		Compare:  s.Compare,
		items: items,
	}
	return clone
}
