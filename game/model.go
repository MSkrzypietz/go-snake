package game

type Model struct {
	Snake Snake
	World World
}

func (m *Model) checkCollision() bool {
	posX, posY := m.Snake.GetPosition()
	return posX < 0 || posY < 0 || posX >= m.World.ColumnCount || posY >= m.World.RowCount
}

func (m *Model) handleFrame() (done bool) {
	m.Snake.Move()
	return m.checkCollision()
}
