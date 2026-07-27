package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// MemoryEntry is a single recorded decision or fact
type MemoryEntry struct {
	category string
	title    string
	detail   string
	tag      string // "decision" | "arch" | "event"
}

// Memory holds the memory panel state
type Memory struct {
	entries []MemoryEntry
	width   int
	height  int
}

// NewMemory creates a memory panel with placeholder data
func NewMemory() Memory {
	return Memory{
		entries: []MemoryEntry{
			{
				category: "Database",
				title:    "CockroachDB over PostgreSQL",
				detail:   "Chosen for horizontal scaling and geo-distribution.",
				tag:      "decision",
			},
			{
				category: "Caching",
				title:    "Redis TTL = 3600s",
				detail:   "Default TTL set after cache-miss incident on 2026-07-24.",
				tag:      "arch",
			},
			{
				category: "Auth",
				title:    "JWT + Refresh tokens",
				detail:   "Access: 24h, Refresh: 30d. Stored in Redis.",
				tag:      "decision",
			},
			{
				category: "Deployment",
				title:    "Docker + Fly.io",
				detail:   "Single region (us-east) for now. Multi-region in Q3.",
				tag:      "arch",
			},
			{
				category: "API",
				title:    "REST + WebSocket",
				detail:   "REST for CRUD, WebSocket for real-time Cortex chat.",
				tag:      "arch",
			},
			{
				category: "Observability",
				title:    "Structured JSON logs",
				detail:   "Using zerolog. Logs shipped to Loki.",
				tag:      "event",
			},
			{
				category: "Security",
				title:    "Rate limiting on /ws",
				detail:   "100 req/min per IP. Implemented in middleware.",
				tag:      "decision",
			},
		},
	}
}

// View renders the memory panel
func (m Memory) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("MEMORY") + "\n\n")

	for _, entry := range m.entries {
		// Category label
		cat := mutedStyle.Render(entry.category)
		tag := renderTag(entry.tag)
		header := lipgloss.JoinHorizontal(lipgloss.Top, cat, "  ", tag)
		b.WriteString(header + "\n")

		// Title
		b.WriteString(itemStyle.Render(entry.title) + "\n")

		// Detail (muted, wrapped)
		detail := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Width(m.width - 6).
			PaddingLeft(1).
			Render(entry.detail)
		b.WriteString(detail + "\n")

		// Separator
		b.WriteString(dividerStyle.Render(strings.Repeat("─", m.width-4)) + "\n")
	}

	return panelStyle.
		Width(m.width).
		Height(m.height).
		Render(b.String())
}

// renderTag returns a styled tag label
func renderTag(tag string) string {
	switch tag {
	case "decision":
		return tagGreen.Render("[decision]")
	case "arch":
		return tagYellow.Render("[arch]")
	case "event":
		return accentStyle.Render("[event]")
	default:
		return mutedStyle.Render("[note]")
	}
}
