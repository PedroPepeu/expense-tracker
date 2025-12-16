// update.go
package ui

import (
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Update

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// FORM
	if m.FormState == StateForm {
		form, cmd := m.form.Update(message)
		if f, ok := form.(*huh.Form); ok {
			m.form = f
			cmds = append(cmds, cmd)

			if m.form.State == huh.StateCompleted {
				if m.formData.Confirm && m.formData.Title != "" {

					spentVal, _ := strconv.Atoi(m.formData.Spent)
					instVal, _ := strconv.Atoi(m.formData.Installment)
					isExpenseBool := m.formData.IsExpense == 1

					newExpense := Expense{
						Title:       m.formData.Title,
						Category:    m.formData.Category,
						Spent:       spentVal,
						Installment: instVal,
						Expense:     isExpenseBool,
						CreatedAt:   time.Now(),
					}

					if m.indexToEdit == -1 {
						// Add
						m.Expenses = append(m.Expenses, newExpense)
					} else {
						// Edit
						m.Expenses[m.indexToEdit] = newExpense
					}
				}

				m.FormState = StateList
				m.formData = &FormData{Installment: "1"}
				m.form = NewExpenseForm(m.formData)
			}
		}
		return m, tea.Batch(cmds...)
	}

	// LIST

	switch msg := message.(type) {
	case tea.ResumeMsg:
		m.suspending = false
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// Global Quits
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "ctrl+z":
			m.suspending = true
			return m, tea.Suspend

			// Alt Screen Toggle
		case " ":
			m.altscreen = !m.altscreen
			return m, nil

			// Switch Focus
		case "tab":
			if m.state == TimerView {
				m.state = spinnerView
			} else {
				m.state = TimerView
			}

		case "n":
			if m.state == TimerView {
				m.timer = timer.New(DefaultTime)
				cmds = append(cmds, m.timer.Init())
			} else {
				m.index = m.NextSpinnerIndex()
				m.spinner = m.GetNewSpinner(m.index)
				cmds = append(cmds, m.spinner.Tick)
			}

			// Open Form
		case "e":
			m.FormState = StateForm
			m.indexToEdit = -1

			m.form = NewExpenseForm(m.formData)
			cmds = append(cmds, m.form.Init())
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)

	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)

	}
	return m, tea.Batch(cmds...)
}
