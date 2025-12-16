package ui

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
)

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
