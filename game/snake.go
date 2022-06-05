package game

type direction int

const (
	North direction = iota
	East
	South
	West
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
		headDirection: South,
	}
}

func (s *Snake) GetPosition() (int, int) {
	return s.head.x, s.head.y
}

func (s *Snake) Move() {
	switch s.headDirection {
	case North:
		s.head.y--
	case East:
		s.head.x++
	case South:
		s.head.y++
	case West:
		s.head.x--
	}
}

func (s *Snake) TurnUp() {
	s.headDirection = North
}

func (s *Snake) TurnRight() {
	s.headDirection = East
}

func (s *Snake) TurnDown() {
	s.headDirection = South
}

func (s *Snake) TurnLeft() {
	s.headDirection = West
}
