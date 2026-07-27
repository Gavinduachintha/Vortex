package main

import (
	"fmt"
	"os"
	// "strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width, height int
	leftFocus     bool // which side is focused
	leftContent   string
	rightContent  string
}

func initialModel() model {
	return model{
		leftFocus:    true,
		leftContent:  "Left pane\n\n• Item 1\n• Item 2\n• Item 3",
		rightContent: "Right pane\n\nThis is the detail / preview side.\nPress Tab to switch focus.",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.leftFocus = !m.leftFocus
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Leave a little room for borders / padding
	leftWidth := m.width / 2
	rightWidth := m.width - leftWidth

	// Styles
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Height(m.height - 2) // leave space for help line

	focusedStyle := borderStyle.Copy().
		BorderForeground(lipgloss.Color("62")) // purple-ish

	unfocusedStyle := borderStyle.Copy().
		BorderForeground(lipgloss.Color("240"))

	var left, right string
	if m.leftFocus {
		left = focusedStyle.Width(leftWidth - 2).Render(m.leftContent)
		right = unfocusedStyle.Width(rightWidth - 2).Render(m.rightContent)
	} else {
		left = unfocusedStyle.Width(leftWidth - 2).Render(m.leftContent)
		right = focusedStyle.Width(rightWidth - 2).Render(m.rightContent)
	}

	// The actual left+right layout
	main := lipgloss.JoinHorizontal(lipgloss.Top, left, right)

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render("\n  Tab: switch focus  •  q: quit")

	return main + help
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}