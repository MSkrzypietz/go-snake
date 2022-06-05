package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"time"

	"go-snake/game"
)

const (
	FPS = 3
)

type frameMsg time.Time

type model struct {
	snake game.Snake
}

func initialModel() model {
	return model{
		snake: game.NewSnake(),
	}
}

func (m model) View() string {
	border := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder(), true, true)

	s := ""
	posX, posY := m.snake.GetPosition()
	for row := 0; row < 20; row++ {
		for col := 0; col < 30; col++ {
			if posX == col && posY == row {
				s += "X"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return border.Render(s)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case frameMsg:
		m.snake.Move()
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
