package utils

import "fmt"

type Point struct {
  X int
  Y int
}

func NewPoint(x int, y int) Point {
  return Point{X: x, Y: y}
}

func (p *Point) String() string {
  return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) Equals(other *Point) bool {
  return p.X == other.X && p.Y == other.Y
}

func (p *Point) Clone() *Point {
  return &Point{X: p.X, Y: p.Y}
}

func (p *Point) Add(other *Point) *Point {
  p.X += other.X
  p.Y += other.Y
  return p
}
