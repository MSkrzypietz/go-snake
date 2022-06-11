package game

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}
