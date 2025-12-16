package ui

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
)

// JSON Functions
func SaveExpensesToJSON(path string, data []Expense) error {
	// 1. Convert struct to JSON bytes
	// Prefix "" and Indent "  " makes it readable (pretty print)
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// 2. Write to file (0644 is standard read/write permission)
	return os.WriteFile(path, file, 0644)
}

// Validators
func ValidateName(str string) error {
	if len(str) < 2 {
		return errors.New("Name must be at least 2 characters.")
	}
	return nil
}

func ValidateAmount(str string) error {
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		return errors.New("Must be a valid number (e.g. 10.50)")
	}
	return nil
}

func ValidateInt(str string) error {
	if _, err := strconv.Atoi(str); err != nil {
		return errors.New("Must be a whole number")
	}
	return nil
}

// Logic helpers

func (m Model) currentFocusedModel() string {
	if m.state == TimerView {
		return "timer"
	}
	return "spinner"
}

func (m Model) NextSpinnerIndex() int {
	if m.index == len(spinners)-1 {
		return 0
	}
	return m.index + 1
}

func (m Model) GetNewSpinner(index int) spinner.Model {
	s := spinner.New()
	s.Style = spinnerStyle
	s.Spinner = spinners[index]
	return s
}

func GetCategoryName(id int) string {
	switch id {
	case 0:
		return "Fun"
	case 1:
		return "Work"
	case 2:
		return "Study"
	case 3:
		return "Investment"
	default:
		return "Other"
	}
}

// Helper to format currency
func FormatAmount(amount int) string {
	return fmt.Sprintf("$ %d.00", amount)
}
