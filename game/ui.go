package game

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	FPS = 3
)

type frameMsg time.Time

func InitialModel() Model {
	return Model{
		Snake: *NewSnake(),
		World: World{
			RowCount:    20,
			ColumnCount: 30,
		},
	}
}

func (m Model) View() string {
	b := strings.Builder{}
	posX, posY := m.Snake.GetPosition()
	for row := 0; row < m.World.RowCount; row++ {
		for col := 0; col < m.World.ColumnCount; col++ {
			if posX == col && posY == row {
				b.WriteRune('X')
			} else {
				b.WriteRune(' ')
			}
		}
		if row < m.World.RowCount-1 {
			b.WriteRune('\n')
		}
	}

	border := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true)
	return border.Render(b.String())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case frameMsg:
		if m.handleFrame() {
			return m, tea.Quit
		}
		return m, animate()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "w":
			m.Snake.TurnUp()
			return m, nil
		case "d":
			m.Snake.TurnRight()
			return m, nil
		case "s":
			m.Snake.TurnDown()
			return m, nil
		case "a":
			m.Snake.TurnLeft()
			return m, nil
		}
	}

	return m, nil
}

func animate() tea.Cmd {
	return tea.Tick(time.Second/FPS, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return animate()
}
