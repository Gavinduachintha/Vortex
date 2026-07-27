package ui

import "github.com/charmbracelet/lipgloss"

// Colors
const (
	colorBackground = "#0D1117"
	colorSurface    = "#161B22"
	colorBorder     = "#30363D"
	colorText       = "#E6EDF3"
	colorMuted      = "#8B949E"
	colorAccent     = "#00E5FF"
	colorAccentDim  = "#0097A7"
	colorGreen      = "#3FB950"
	colorYellow     = "#D29922"
	colorRed        = "#F85149"
)

// Panel widths
const (
	sidebarWidth = 24
	memoryWidth  = 32
)

// Shared border style
var panelBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "╰",
	BottomRight: "╯",
}

// Base panel style — width is set per-component
var panelStyle = lipgloss.NewStyle().
	Border(panelBorder).
	BorderForeground(lipgloss.Color(colorBorder)).
	Padding(0, 1)

// Title bar inside a panel
var titleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorAccent)).
	Bold(true).
	MarginBottom(1)

// Normal item in a list
var itemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorText)).
	PaddingLeft(1)

// Selected item in a list
var selectedItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorAccent)).
	Bold(true).
	PaddingLeft(1)

// Muted / secondary text
var mutedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorMuted))

// Accent text
var accentStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorAccent))

// Divider line
var dividerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorBorder))

// Input prompt
var inputPromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorAccent)).
	Bold(true)

// Status tag styles
var tagGreen = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorGreen)).
	Bold(true)

var tagYellow = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorYellow)).
	Bold(true)

var tagRed = lipgloss.NewStyle().
	Foreground(lipgloss.Color(colorRed)).
	Bold(true)
