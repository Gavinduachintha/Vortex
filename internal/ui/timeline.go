package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TimelineEvent is a single project history entry
type TimelineEvent struct {
	date    string
	title   string
	detail  string
	kind    string // "release" | "incident" | "decision" | "milestone"
}

// Timeline holds project history state
type Timeline struct {
	events []TimelineEvent
	width  int
	height int
}

// NewTimeline creates a timeline panel with placeholder history
func NewTimeline() Timeline {
	return Timeline{
		events: []TimelineEvent{
			{
				date:   "2026-07-25",
				title:  "v0.4.2 Released",
				detail: "WebSocket stability fixes. Cortex chat latency reduced 40%.",
				kind:   "release",
			},
			{
				date:   "2026-07-24",
				title:  "Incident: Redis cache miss spike",
				detail: "TTL misconfiguration caused 8% error rate. Resolved in 40 min.",
				kind:   "incident",
			},
			{
				date:   "2026-07-22",
				title:  "Memory panel designed",
				detail: "Decided to surface decisions, arch choices, and events in TUI.",
				kind:   "decision",
			},
			{
				date:   "2026-07-20",
				title:  "v0.4.0 Released",
				detail: "Initial TUI shipped. Sidebar + Chat functional.",
				kind:   "release",
			},
			{
				date:   "2026-07-15",
				title:  "CockroachDB migration complete",
				detail: "Full cutover from PostgreSQL. Zero downtime migration.",
				kind:   "milestone",
			},
			{
				date:   "2026-07-10",
				title:  "Fly.io deployment live",
				detail: "Production environment up. us-east region.",
				kind:   "milestone",
			},
			{
				date:   "2026-07-05",
				title:  "JWT auth implemented",
				detail: "Access + refresh token flow. Integrated with Redis.",
				kind:   "decision",
			},
			{
				date:   "2026-07-01",
				title:  "Project initialized",
				detail: "Cortex TUI project scaffolded. Initial module structure set.",
				kind:   "milestone",
			},
		},
	}
}

// View renders the timeline panel (full-width, used when timeline view is active)
func (t Timeline) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("TIMELINE") + "\n\n")

	for i, event := range t.events {
		// Connector line
		connector := "│"
		if i == 0 {
			connector = "◉"
		}

		// Date + kind badge
		date := mutedStyle.Render(event.date)
		badge := renderTimelineBadge(event.kind)
		header := lipgloss.JoinHorizontal(lipgloss.Top,
			accentStyle.Render(connector+"  "),
			date,
			"  ",
			badge,
		)
		b.WriteString(header + "\n")

		// Title
		titleLine := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorText)).
			Bold(true).
			PaddingLeft(4).
			Render(event.title)
		b.WriteString(titleLine + "\n")

		// Detail
		detail := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Width(t.width - 8).
			PaddingLeft(4).
			Render(event.detail)
		b.WriteString(detail + "\n")

		// Spacer (except last)
		if i < len(t.events)-1 {
			b.WriteString(accentStyle.Render("│") + "\n")
		}
	}

	return panelStyle.
		Width(t.width).
		Height(t.height).
		Render(b.String())
}

// renderTimelineBadge returns a color-coded badge for the event kind
func renderTimelineBadge(kind string) string {
	switch kind {
	case "release":
		return tagGreen.Render("[release]")
	case "incident":
		return tagRed.Render("[incident]")
	case "decision":
		return tagYellow.Render("[decision]")
	case "milestone":
		return accentStyle.Render("[milestone]")
	default:
		return mutedStyle.Render("[event]")
	}
}
