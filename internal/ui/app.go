package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// activeView represents which main panel is shown in the center
type activeView int

const (
	viewChat activeView = iota
	viewMemory
	viewTimeline
)

// App is the root Bubble Tea model
type App struct {
	sidebar    Sidebar
	chat       Chat
	memory     Memory
	timeline   Timeline
	active     activeView
	width      int
	height     int
	sidebarFoc bool // whether sidebar has focus
}

// NewApp initializes the application
func NewApp() App {
	return App{
		sidebar:    NewSidebar(),
		chat:       NewChat(),
		memory:     NewMemory(),
		timeline:   NewTimeline(),
		active:     viewChat,
		sidebarFoc: false,
	}
}

// Init is called once on startup; no initial commands needed
func (a App) Init() tea.Cmd {
	return nil
}

// Update handles all incoming messages
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.recalcSizes()
		return a, nil

	case tea.KeyMsg:
		switch msg.String() {

		// Global quit
		case "ctrl+q":
			return a, tea.Quit

		// Switch to chat view, focus chat
		case "ctrl+c":
			a.active = viewChat
			a.sidebarFoc = false
			return a, nil

		// Switch to memory view
		case "ctrl+m":
			a.active = viewMemory
			a.sidebarFoc = false
			return a, nil

		// Switch to timeline view
		case "ctrl+t":
			a.active = viewTimeline
			a.sidebarFoc = false
			return a, nil

		// Toggle sidebar focus
		case "tab":
			a.sidebarFoc = !a.sidebarFoc
			return a, nil

		// Arrow / vim keys — route to sidebar or chat
		case "up", "down", "ctrl+k", "ctrl+j":
			if a.sidebarFoc {
				var selectedView string
				a.sidebar, selectedView = a.sidebar.Update(msg)
				if selectedView != "" {
					a.active = stringToView(selectedView)
				}
			} else if a.active == viewChat {
				// No scroll yet — placeholder
			}
			return a, nil

		// Enter — confirm sidebar selection or send chat message
		case "enter":
			if a.sidebarFoc {
				var selectedView string
				a.sidebar, selectedView = a.sidebar.Update(msg)
				if selectedView != "" {
					a.active = stringToView(selectedView)
					a.sidebarFoc = false
				}
			} else if a.active == viewChat {
				a.chat = a.chat.Update(msg)
			}
			return a, nil

		// All other keys go to chat input when chat is active
		default:
			if a.active == viewChat && !a.sidebarFoc {
				a.chat = a.chat.Update(msg)
			}
		}
	}

	return a, nil
}

// View renders the full terminal layout
func (a App) View() string {
	if a.width == 0 {
		return "Loading…"
	}

	sidebar := a.sidebar.View()
	center := a.centerView()
	right := a.rightView()

	body := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, center, right)

	// Help bar at the bottom
	help := a.helpBar()

	return lipgloss.JoinVertical(lipgloss.Left, body, help)
}

// centerView returns the active center panel
func (a App) centerView() string {
	switch a.active {
	case viewMemory:
		return a.memory.View()
	case viewTimeline:
		return a.timeline.View()
	default:
		return a.chat.View()
	}
}

// rightView returns the right panel (always memory summary or timeline)
func (a App) rightView() string {
	switch a.active {
	case viewTimeline:
		// When timeline is center, show memory in the right column
		return a.memory.View()
	default:
		// Otherwise right panel always shows timeline summary
		return a.timeline.View()
	}
}

// recalcSizes distributes terminal dimensions to each component
func (a *App) recalcSizes() {
	// -2 per panel for borders, -1 for help bar
	availH := a.height - 1

	a.sidebar.height = availH - 2 // border overhead

	// Center takes remaining width after sidebar + right + their borders
	// Each panel border = 2 chars (left+right)
	rightW := memoryWidth
	centerW := a.width - sidebarWidth - rightW - 6 // 3 panels × 2 border chars
	if centerW < 20 {
		centerW = 20
	}

	a.chat.width = centerW
	a.chat.height = availH - 2

	a.memory.width = rightW
	a.memory.height = availH - 2

	a.timeline.width = rightW
	a.timeline.height = availH - 2
}

// helpBar renders the keyboard shortcut hint strip
func (a App) helpBar() string {
	keys := []string{
		keyHint("ctrl+c", "chat"),
		keyHint("ctrl+m", "memory"),
		keyHint("ctrl+t", "timeline"),
		keyHint("tab", "focus sidebar"),
		keyHint("↑↓", "navigate"),
		keyHint("enter", "select"),
		keyHint("ctrl+q", "quit"),
	}

	bar := strings.Join(keys, mutedStyle.Render("  ·  "))

	return lipgloss.NewStyle().
		Width(a.width).
		Foreground(lipgloss.Color(colorMuted)).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(colorBorder)).
		Padding(0, 1).
		Render(bar)
}

// keyHint formats a single key binding hint
func keyHint(key, desc string) string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorAccent)).
		Bold(true).
		Render(key)
	d := mutedStyle.Render(" " + desc)
	return k + d
}

// stringToView converts a view name string to an activeView constant
func stringToView(v string) activeView {
	switch v {
	case "memory":
		return viewMemory
	case "timeline":
		return viewTimeline
	default:
		return viewChat
	}
}
