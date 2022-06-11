package game

type direction int

const (
	up direction = iota
	right
	down
	left
)

type Snake struct {
	head          *Point
	headDirection direction
	tail          []Point
	tailLength    int
}

func NewSnake() *Snake {
	return &Snake{
		head:          NewPoint(10, 10),
		headDirection: up,
		tail:          []Point{*NewPoint(10, 11), *NewPoint(10, 12)},
		tailLength:    2,
	}
}

func (s *Snake) CollidesWithPoint(p *Point) bool {
	return s.head.Equals(p) || s.TailCollidesWithPoint(p)
}

func (s *Snake) TailCollidesWithPoint(p *Point) bool {
	for _, tailPoint := range s.tail {
		if tailPoint.Equals(p) {
			return true
		}
	}
	return false

}

func (s *Snake) Move() {
	s.tail = append([]Point{*NewPoint(s.head.X, s.head.Y)}, s.tail...)
	s.tail = s.tail[:s.tailLength]

	switch s.headDirection {
	case up:
		s.head.Y--
	case right:
		s.head.X++
	case down:
		s.head.Y++
	case left:
		s.head.X--
	}
}

func (s *Snake) Eat() {
	s.tailLength++
}

func (s *Snake) TurnUp() {
	if s.headDirection != down {
		s.headDirection = up
	}
}

func (s *Snake) TurnRight() {
	if s.headDirection != left {
		s.headDirection = right
	}
}

func (s *Snake) TurnDown() {
	if s.headDirection != up {
		s.headDirection = down
	}
}

func (s *Snake) TurnLeft() {
	if s.headDirection != right {
		s.headDirection = left
	}
}
