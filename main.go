package main

// A simple program that opens the alternate screen buffer then counts down
// from 5 and then exits.

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"expense-tracker/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	storagePath := "store/data.json"

	expenses, err := LoadExpensesFromJSON(storagePath)
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		os.Exit(1)
	}

	m := ui.InitialModel(ui.DefaultTime)
	m.Expenses = expenses

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func LoadExpensesFromJSON(path string) ([]ui.Expense, error) {
	// A. Check if file exists. If not, return empty list (not an error)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []ui.Expense{}, nil
	}

	// B. Open the file
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	// C. Check if file is empty
	if len(file) == 0 {
		return []ui.Expense{}, nil
	}

	// D. Unmarshal (Parse) JSON into the Slice
	var data []ui.Expense
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("could not parse json: %w", err)
	}

	return data, nil
}
