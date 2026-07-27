package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Message represents a single chat entry
type Message struct {
	role    string // "user" or "cortex"
	content string
}

// Chat holds conversation state and input
type Chat struct {
	messages []Message
	input    string
	width    int
	height   int
}

// NewChat initializes the chat panel with some seed messages
func NewChat() Chat {
	return Chat{
		messages: []Message{
			{role: "cortex", content: "Cortex online. How can I help?"},
			{role: "user", content: "What databases are we using?"},
			{role: "cortex", content: "CockroachDB for primary storage, Redis for caching.\nMigration from PostgreSQL completed last sprint."},
			{role: "user", content: "Any incidents this week?"},
			{role: "cortex", content: "One incident: Redis cache miss spike on 2026-07-24.\nRoot cause: TTL misconfiguration. Resolved in ~40 min."},
		},
		input: "",
	}
}

// Update handles key messages for the chat panel
func (c Chat) Update(msg tea.KeyMsg) Chat {
	switch msg.String() {
	case "enter":
		if strings.TrimSpace(c.input) != "" {
			userMsg := Message{role: "user", content: c.input}
			fakeReply := Message{role: "cortex", content: fakeResponse(c.input)}
			c.messages = append(c.messages, userMsg, fakeReply)
			c.input = ""
		}
	case "backspace":
		if len(c.input) > 0 {
			c.input = c.input[:len(c.input)-1]
		}
	default:
		// Only append printable single-char keys
		if len(msg.String()) == 1 {
			c.input += msg.String()
		}
	}
	return c
}

// View renders the chat panel
func (c Chat) View() string {
	// Reserve space: borders (2) + padding (2) + input area (3)
	const inputAreaHeight = 3
	contentHeight := c.height - inputAreaHeight - 4
	if contentHeight < 1 {
		contentHeight = 1
	}

	// Build message list, keep last N that fit
	var lines []string
	for _, m := range c.messages {
		lines = append(lines, renderMessage(m, c.width-4))
	}

	// Trim to visible height (simple line-based crop from bottom)
	var visible []string
	totalLines := 0
	for i := len(lines) - 1; i >= 0; i-- {
		msgLines := strings.Count(lines[i], "\n") + 2 // +1 for prefix, +1 for spacing
		if totalLines+msgLines > contentHeight {
			break
		}
		visible = append([]string{lines[i]}, visible...)
		totalLines += msgLines
	}

	// Pad top so messages sit at bottom
	padding := contentHeight - totalLines
	if padding < 0 {
		padding = 0
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("CHAT") + "\n")
	b.WriteString(strings.Repeat("\n", padding))
	b.WriteString(strings.Join(visible, "\n"))

	// Input divider
	divider := dividerStyle.Render(strings.Repeat("─", c.width-4))
	b.WriteString("\n" + divider + "\n")

	// Input line
	prompt := inputPromptStyle.Render("▸ ")
	cursor := accentStyle.Render("█")
	b.WriteString(prompt + c.input + cursor)

	return panelStyle.
		Width(c.width).
		Height(c.height).
		Render(b.String())
}

// renderMessage formats a single message with role prefix
func renderMessage(m Message, width int) string {
	switch m.role {
	case "user":
		prefix := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorText)).
			Bold(true).
			Render("You  ")
		body := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorText)).
			Width(width - 5).
			Render(m.content)
		return prefix + body
	default: // cortex
		prefix := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorAccent)).
			Bold(true).
			Render("Ctx  ")
		body := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Width(width - 5).
			Render(m.content)
		return prefix + body
	}
}

// fakeResponse returns a canned reply for demo purposes
func fakeResponse(input string) string {
	input = strings.ToLower(input)
	switch {
	case strings.Contains(input, "database") || strings.Contains(input, "db"):
		return "Using CockroachDB (primary) and Redis (cache). No schema changes pending."
	case strings.Contains(input, "auth") || strings.Contains(input, "jwt"):
		return "JWT-based authentication. Tokens expire in 24h. Refresh tokens stored in Redis."
	case strings.Contains(input, "incident") || strings.Contains(input, "error"):
		return "Last incident: Redis TTL misconfiguration on 2026-07-24. All systems nominal now."
	case strings.Contains(input, "deploy") || strings.Contains(input, "release"):
		return "Last deploy: v0.4.2 on 2026-07-25. Next release scheduled for 2026-08-01."
	default:
		return "Logged. I'll surface relevant context as the project evolves."
	}
}
