package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

// Initial Model
func InitialModel(timeout time.Duration) Model {
	data := &FormData{
		Installment: "1", // default
	}

	return Model{
		// Expenses
		Expenses: []Expense{},

		// Composable
		state:   TimerView,
		timer:   timer.New(timeout),
		spinner: spinner.New(),
		index:   0,

		// Form
		FormState:   StateList,
		formData:    data,
		form:        NewExpenseForm(data),
		indexToEdit: -1,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.timer.Init(), m.spinner.Tick)
}
