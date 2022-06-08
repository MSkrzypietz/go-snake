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
}

func NewSnake() *Snake {
	return &Snake{
		head:          NewPoint(10, 10),
		headDirection: right,
	}
}

func (s *Snake) GetPosition() (int, int) {
	return s.head.X, s.head.Y
}

func (s *Snake) Move() {
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
