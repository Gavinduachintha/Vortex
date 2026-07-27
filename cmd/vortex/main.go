package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"Vertex/internal/ui"
)

func main() {
	app := ui.NewApp()

	p := tea.NewProgram(
		app,
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
