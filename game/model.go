package game

import "math/rand"

type Model struct {
	Snake Snake
	World World
	Food  *Point
}

func (m *Model) handleFrame() (done bool) {
	m.Snake.Move()
	if m.checkFoodCollision() {
		m.Snake.Eat()
		m.respawnFood()
	}
	return m.checkBorderCollision() || m.Snake.TailCollidesWithPoint(m.Snake.head)
}

func (m *Model) checkFoodCollision() bool {
	return m.Snake.CollidesWithPoint(m.Food)
}

func (m *Model) checkBorderCollision() bool {
	posX, posY := m.Snake.head.X, m.Snake.head.Y
	return posX < 0 || posY < 0 || posX >= m.World.ColumnCount || posY >= m.World.RowCount
}

func (m *Model) respawnFood() {
	m.Food = NewPoint(rand.Intn(m.World.ColumnCount), rand.Intn(m.World.RowCount))
}
