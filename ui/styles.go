package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Lipgloss
	spinners = []spinner.Spinner{
		spinner.Line, spinner.Dot, spinner.MiniDot, spinner.Jump,
		spinner.Pulse, spinner.Points, spinner.Globe, spinner.Moon, spinner.Monkey,
	}
	modelStyle = lipgloss.NewStyle().
			Width(15).Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())

	focusedModelStyle = lipgloss.NewStyle().
				Width(15).Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))

	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("204")).Background(lipgloss.Color("235"))

	listHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("252")).
			Border(lipgloss.NormalBorder(), false, false, true, false). // Bottom border
			Padding(0, 1)

	listItemStyle = lipgloss.NewStyle().PaddingLeft(2)

	profitStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))  // Green
	lossStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red
	dimmedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Grey for category
)
