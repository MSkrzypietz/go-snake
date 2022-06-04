package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

const (
	FPS = 3
)

type frameMsg time.Time

type Direction int

const (
	North Direction = iota
	West
	South
	East
)

type head struct {
	x int
	y int
}

type snake struct {
	head          head
	currDirection Direction
}

func (s *snake) move() {
	switch s.currDirection {
	case North:
		s.head.y--
	case West:
		s.head.x++
	case South:
		s.head.y++
	case East:
		s.head.x--
	}
}

type model struct {
	snake snake
}

func initialModel() model {
	return model{
		snake: snake{
			head:          head{x: 2, y: 2},
			currDirection: South,
		},
	}
}

func (m model) View() string {
	s := ""
	for row := 0; row < 30; row++ {
		for col := 0; col < 10; col++ {
			if m.snake.head.x == col && m.snake.head.y == row {
				s += "X"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case frameMsg:
		m.snake.move()
		return m, animate()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "w":
			m.snake.currDirection = North
			return m, nil
		case "d":
			m.snake.currDirection = West
			return m, nil
		case "s":
			m.snake.currDirection = South
			return m, nil
		case "a":
			m.snake.currDirection = East
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
