package utils

import "fmt"

type Map2d[T comparable] struct {
  Width int;
  Height int;
  Map [][]T
}

func CreateMap2d[T comparable](Map [][]T) *Map2d[T] {
  return &Map2d[T]{
    Width: len(Map[0]),
    Height: len(Map),
    Map: Map,
  }
}

func (Map *Map2d[T]) Get(p *Point) (*T, bool) {
  if p.X < 0 || p.X >= Map.Width || p.Y < 0 || p.Y >= Map.Height {
    return nil, false
  }

  return &Map.Map[p.Y][p.X], true
}

func (Map *Map2d[T]) Set(value T, p *Point) {
  x, y := p.X, p.Y
  Map.Map[y][x] = value
}

func (Map *Map2d[T]) FindPosition(value T) *Point {
  for y := 0; y < Map.Height; y++ {
    for x := 0; x < Map.Width; x++ {
      if Map.Map[y][x] == value {
        return &Point{X: x, Y: y}
      }
    }
  }
  return nil
}

func (Map *Map2d[T]) Print() {
  fmt.Printf("Map %d x %d:\n", Map.Width, Map.Height)
  for _, row := range Map.Map {
    for _, value := range row {
      fmt.Printf("%v ", value)
    }
    fmt.Println()
  }
}

func (Map *Map2d[T]) Clone() *Map2d[T] {
  map2d := make([][]T, Map.Height)
  for y := 0; y < Map.Height; y++ {
    map2d[y] = make([]T, Map.Width)
    copy(map2d[y], Map.Map[y])
  }
  return &Map2d[T]{
    Width: Map.Width,
    Height: Map.Height,
    Map: map2d,
  }
}
