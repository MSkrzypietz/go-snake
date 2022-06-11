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
		headDirection: right,
		tail:          []Point{},
		tailLength:    0,
	}
}

func (s *Snake) CollidesWithPoint(p *Point) bool {
	if s.head.Equals(p) {
		return true
	}

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
	s.headDirection = up
}

func (s *Snake) TurnRight() {
	s.headDirection = right
}

func (s *Snake) TurnDown() {
	s.headDirection = down
}

func (s *Snake) TurnLeft() {
	s.headDirection = left
}
