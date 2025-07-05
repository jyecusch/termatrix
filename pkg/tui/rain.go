package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jyecusch/termatrix/pkg/colors"
	"github.com/jyecusch/termatrix/pkg/matrix"
)

type matrixRain struct {
	grid       *matrix.RainGrid
	timer      timer.Model
	ignoreTick bool
	colorIndex int
}

func (m matrixRain) Init() tea.Cmd {
	return m.timer.Init()
}

func (m matrixRain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.grid.Resize(msg.Width, msg.Height)
		return m, nil
	case timer.TickMsg:
		if !m.ignoreTick {
			m.grid.Update()
		}
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "Q", "ctrl+c":
			return m, tea.Quit
		case "p":
			m.ignoreTick = !m.ignoreTick
			return m, nil
		case "t":
			m.grid.Update()
			return m, nil
		case "c":
			m.colorIndex = (m.colorIndex + 1) % len(colors.AllColors)
			return m, nil
		default:
			return m, nil
		}
	}
	return m, nil
}

func (m matrixRain) View() string {
	return m.grid.Draw(colors.AllColors[m.colorIndex])
}

func NewMatrixRain() matrixRain {
	return matrixRain{
		grid: matrix.NewRainGrid(100, 100),
		// TODO: replace with a custom ticker instead of timer
		timer:      timer.NewWithInterval(1000*time.Hour, time.Second/22),
		colorIndex: 0,
	}
}
