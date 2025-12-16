package ui

import (
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// 1. LÓGICA DO FORMULÁRIO (Se estiver aberto)
	if m.FormState == StateForm {
		form, cmd := m.form.Update(message)
		if f, ok := form.(*huh.Form); ok {
			m.form = f
			cmds = append(cmds, cmd)

			if m.form.State == huh.StateCompleted {
				if m.formData.Confirm && m.formData.Title != "" {
					// Conversões
					spentVal, _ := strconv.Atoi(m.formData.Spent)
					instVal, _ := strconv.Atoi(m.formData.Installment)
					isExpenseBool := m.formData.IsExpense == 1
					parsedDate, err := time.Parse("02-01-2006", m.formData.Date)
					if err != nil {
						parsedDate = time.Now()
					}

					newExpense := Expense{
						// ID simples baseado no tamanho da lista
						ID:          len(m.Expenses) + 1,
						Title:       m.formData.Title,
						Category:    m.formData.Category,
						Spent:       spentVal,
						Installment: instVal,
						Expense:     isExpenseBool,
						Date:        parsedDate,
						CreatedAt:   time.Now(),
					}

					if m.indexToEdit == -1 {
						m.Expenses = append(m.Expenses, newExpense)
					} else {
						m.Expenses[m.indexToEdit] = newExpense
					}

					// Salvar
					_ = SaveExpensesToJSON("store/data.json", m.Expenses)
				}

				// Resetar
				m.FormState = StateList
				m.formData = &FormData{Installment: "1"}
				m.form = NewExpenseForm(m.formData)
				m.indexToEdit = -1
			}
		}
		return m, tea.Batch(cmds...)
	}

	// 2. LÓGICA DA APLICAÇÃO PRINCIPAL
	switch msg := message.(type) {
	case tea.ResumeMsg:
		m.suspending = false
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// --- COMANDOS GERAIS ---
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "ctrl+z":
			m.suspending = true
			return m, tea.Suspend

		case " ":
			m.altscreen = !m.altscreen
			return m, nil

		// --- COMANDOS DE NAVEGAÇÃO (LISTA) ---
		// AQUI ESTÁ A CORREÇÃO QUE FALTAVA:
		case "up", "k":
			if m.altscreen && m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.altscreen && m.Cursor < len(m.Expenses)-1 {
				m.Cursor++
			}

		// --- COMANDO DE DELETAR ---
		case "d":
			if m.altscreen && len(m.Expenses) > 0 {
				// Remove o item da lista
				m.Expenses = append(m.Expenses[:m.Cursor], m.Expenses[m.Cursor+1:]...)

				// Ajusta o cursor para não ficar fora da lista
				if m.Cursor >= len(m.Expenses) && len(m.Expenses) > 0 {
					m.Cursor = len(m.Expenses) - 1
				} else if len(m.Expenses) == 0 {
					m.Cursor = 0
				}
				// Salva
				_ = SaveExpensesToJSON("store/data.json", m.Expenses)
			}

		// --- COMANDOS DO DASHBOARD (Inline) ---
		case "tab":
			if !m.altscreen {
				if m.state == TimerView {
					m.state = SpinnerView
				} else {
					m.state = TimerView
				}
			}

		case "n":
			if !m.altscreen {
				if m.state == TimerView {
					m.timer = timer.New(DefaultTime)
					cmds = append(cmds, m.timer.Init())
				} else {
					m.index = m.NextSpinnerIndex()
					m.spinner = m.GetNewSpinner(m.index)
					cmds = append(cmds, m.spinner.Tick)
				}
			}

		// --- ABRIR FORMULÁRIO ---
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
