package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SidebarItem represents a single navigation entry
type SidebarItem struct {
	label string
	icon  string
	view  string // the view key this item maps to
}

// Sidebar holds navigation state
type Sidebar struct {
	items    []SidebarItem
	selected int
	height   int
}

// NewSidebar creates a sidebar with default items
func NewSidebar() Sidebar {
	return Sidebar{
		items: []SidebarItem{
			{icon: "󰉋", label: "Files", view: "chat"},
			{icon: "󰙶", label: "Memories", view: "memory"},
			{icon: "󰙭", label: "Decisions", view: "memory"},
			{icon: "󰃤", label: "Incidents", view: "timeline"},
			{icon: "󰔖", label: "Timeline", view: "timeline"},
		},
		selected: 0,
	}
}

// Update handles key messages for sidebar navigation
func (s Sidebar) Update(msg tea.KeyMsg) (Sidebar, string) {
	switch msg.String() {
	case "up", "k":
		if s.selected > 0 {
			s.selected--
		}
	case "down", "j":
		if s.selected < len(s.items)-1 {
			s.selected++
		}
	case "enter":
		return s, s.items[s.selected].view
	}
	return s, ""
}

// View renders the sidebar panel
func (s Sidebar) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("CORTEX") + "\n\n")

	for i, item := range s.items {
		label := item.icon + "  " + item.label

		if i == s.selected {
			// Highlight bar using a background strip
			row := lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorAccent)).
				Bold(true).
				PaddingLeft(1).
				Render("▸ " + label)
			b.WriteString(row)
		} else {
			row := lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorMuted)).
				PaddingLeft(1).
				Render("  " + label)
			b.WriteString(row)
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(dividerStyle.Render(strings.Repeat("─", sidebarWidth-4)) + "\n\n")

	// Status section at bottom
	b.WriteString(mutedStyle.Render("  Status") + "\n")
	b.WriteString("  " + tagGreen.Render("● ") + mutedStyle.Render("Connected") + "\n")
	b.WriteString("  " + tagYellow.Render("● ") + mutedStyle.Render("3 Pending") + "\n")

	return panelStyle.
		Width(sidebarWidth).
		Height(s.height).
		Render(b.String())
}
