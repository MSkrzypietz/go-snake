package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"go-snake/game"
)

const (
	FPS = 3
)

type frameMsg time.Time

type model struct {
	snake game.Snake
	world game.World
}

func (m model) checkCollision() bool {
	posX, posY := m.snake.GetPosition()
	return posX < 0 || posY < 0 || posX >= m.world.ColumnCount || posY >= m.world.RowCount
}

func initialModel() model {
	return model{
		snake: game.NewSnake(),
		world: game.World{
			RowCount:    20,
			ColumnCount: 30,
		},
	}
}

func (m model) View() string {
	b := strings.Builder{}
	posX, posY := m.snake.GetPosition()
	for row := 0; row < m.world.RowCount; row++ {
		for col := 0; col < m.world.ColumnCount; col++ {
			if posX == col && posY == row {
				b.WriteRune('X')
			} else {
				b.WriteRune(' ')
			}
		}
		if row < m.world.RowCount-1 {
			b.WriteRune('\n')
		}
	}

	border := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true)
	return border.Render(b.String())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case frameMsg:
		m.snake.Move()
		if m.checkCollision() {
			return m, tea.Quit
		}
		return m, animate()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "w":
			m.snake.TurnUp()
			return m, nil
		case "d":
			m.snake.TurnRight()
			return m, nil
		case "s":
			m.snake.TurnDown()
			return m, nil
		case "a":
			m.snake.TurnLeft()
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

func (m model) Init() tea.Cmd {
	return animate()
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
