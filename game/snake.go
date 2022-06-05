package game

type direction int

const (
	up direction = iota
	right
	down
	left
)

type head struct {
	x int
	y int
}

type Snake struct {
	head          head
	headDirection direction
}

func NewSnake() Snake {
	return Snake{
		head:          head{x: 2, y: 2},
		headDirection: down,
	}
}

func (s *Snake) GetPosition() (int, int) {
	return s.head.x, s.head.y
}

func (s *Snake) Move() {
	switch s.headDirection {
	case up:
		s.head.y--
	case right:
		s.head.x++
	case down:
		s.head.y++
	case left:
		s.head.x--
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
