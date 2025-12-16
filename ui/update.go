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

			if m.form.State == huh.StateCompleted {
				// Only save if Confirmed and Title is not empty
				if m.formData.Confirm && m.formData.Title != "" {

					// 1. CONVERT DATA (String -> Int/Time)
					spentVal, _ := strconv.Atoi(m.formData.Spent)
					instVal, _ := strconv.Atoi(m.formData.Installment)
					isExpenseBool := m.formData.IsExpense == 1

					// Parse Date (default to Now if empty/invalid)
					parsedDate, err := time.Parse("02-01-2006", m.formData.Date)
					if err != nil {
						parsedDate = time.Now()
					}

					// 2. LOGIC: NEW vs EDIT
					if m.indexToEdit == -1 {
						// Create NEW Expense
						newExpense := Expense{
							// Auto-increment ID based on list length
							ID:          len(m.Expenses) + 1,
							Title:       m.formData.Title,
							Category:    m.formData.Category,
							Spent:       spentVal,
							Installment: instVal,
							Expense:     isExpenseBool,
							Date:        parsedDate,
							CreatedAt:   time.Now(),
						}
						m.Expenses = append(m.Expenses, newExpense)

					} else {
						// EDIT Existing Expense
						m.Expenses[m.indexToEdit].Title = m.formData.Title
						m.Expenses[m.indexToEdit].Category = m.formData.Category
						m.Expenses[m.indexToEdit].Spent = spentVal
						m.Expenses[m.indexToEdit].Installment = instVal
						m.Expenses[m.indexToEdit].Expense = isExpenseBool
						m.Expenses[m.indexToEdit].Date = parsedDate
						// We don't update CreatedAt
					}

					// 3. SAVE TO JSON
					// Make sure "store" folder exists in your project root!
					if err := SaveExpensesToJSON("store/data.json", m.Expenses); err != nil {
						// If you have a status message field, set it here.
						// For now, we just print to console (which might be hidden by TUI)
						// log.Printf("Error saving: %v", err)
					}
				}

				// 4. RESET FORM
				m.FormState = StateList
				m.formData = &FormData{Installment: "1"} // Reset data
				m.form = NewExpenseForm(m.formData)      // Recreate form
				m.indexToEdit = -1
			}
		}
		return m, cmd
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
