package main

// A simple program that opens the alternate screen buffer then counts down
// from 5 and then exits.

import (
	"log"

	"expense-tracker/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialModel(ui.DefaultTime), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
