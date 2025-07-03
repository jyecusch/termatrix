package matrix

import (
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/jyecusch/termatrix/pkg/colors"
)

func CalcStatus(line RainLine, charPos int) RuneStatus {
	if line.isHead(charPos) {
		return RuneStatusHead
	}

	if line.isTail(charPos) {
		return RuneStatusTail
	}

	if line.isInside(charPos) {
		return RuneStatusActive
	}

	return RuneStatusInactive
}

type RainGrid struct {
	width   int
	height  int
	columns []*RainLine
	lock    sync.Mutex

	grid [][]rune

	intro     string
	introTime time.Duration
	start     time.Time
}

func NewRainGrid(width, height int) *RainGrid {
	columns := make([]*RainLine, width)
	for col := range columns {
		columns[col] = NewRandomLine(height)
	}

	return &RainGrid{
		width:     width,
		height:    height,
		columns:   columns,
		grid:      newRuneGrid(width, height),
		intro:     "The Matrix Has You",
		introTime: time.Second * 2,
		start:     time.Now(),
	}
}

func (r *RainGrid) Update() {
	r.lock.Lock()
	defer r.lock.Unlock()

	for col := range r.columns {
		r.columns[col].update(r.height, time.Now())
	}
}

func (r *RainGrid) Resize(width, height int) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if width == r.width && height == r.height {
		return
	}

	r.grid = newRuneGrid(width, height)
	r.columns = make([]*RainLine, width)

	for col := range r.columns {
		r.columns[col] = NewRandomLine(height)
	}

	r.height = height
	r.width = width
}

func statusStyles(color lipgloss.Color) map[RuneStatus]lipgloss.Style {
	return map[RuneStatus]lipgloss.Style{
		RuneStatusHead:     lipgloss.NewStyle().Foreground(colors.White).Bold(true),
		RuneStatusTail:     lipgloss.NewStyle().Foreground(color).Faint(true),
		RuneStatusActive:   lipgloss.NewStyle().Foreground(color),
		RuneStatusInactive: lipgloss.NewStyle(),
	}
}

func (r *RainGrid) Draw(color lipgloss.Color) string {
	r.lock.Lock()
	defer r.lock.Unlock()

	statusStyles := statusStyles(color)

	var output strings.Builder

	now := time.Now()

	for row := range r.height {
		isCenterRow := row == r.height/2
		introStartIndex := r.width/2 - len(r.intro)/2

		for col := range r.width {
			if isCenterRow && col >= introStartIndex && col < introStartIndex+len(r.intro) {
				charIx := col - introStartIndex

				fadeTicks := 4
				timePerChar := r.introTime / time.Duration(len(r.intro)+fadeTicks)
				fadeTime := time.Duration(timePerChar * time.Duration(charIx+fadeTicks/2))
				fullTime := time.Duration(timePerChar * time.Duration(charIx+fadeTicks))
				shouldFade := now.Sub(r.start) > fadeTime
				shouldSkip := now.Sub(r.start) > fullTime

				style := statusStyles[RuneStatusActive]
				if shouldFade {
					style = style.Faint(true)
				}

				if !shouldSkip {
					output.WriteString(style.Render(string(r.intro[charIx])))
					continue
				}
			}

			status := CalcStatus(*r.columns[col], row)

			if status == RuneStatusInactive {
				r.grid[col][row] = ' '
			} else if r.grid[col][row] == ' ' {
				r.grid[col][row] = RandomRune()
			}

			output.WriteString(statusStyles[status].Render(string(r.grid[col][row])))
		}
		output.WriteString("\n")
	}

	return output.String()
}
