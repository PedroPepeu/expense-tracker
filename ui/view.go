// view.go
package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.suspending {
		return ""
	}

	if m.quitting {
		return "Bye!\n"
	}

	if m.FormState == StateForm {
		return m.form.View()
	}

	var content string

	var timerStr, spinnerStr string
	if m.state == TimerView {
		timerStr = focusedModelStyle.Render(fmt.Sprintf("%4s", m.timer.View()))
		spinnerStr = modelStyle.Render(m.spinner.View())
	} else {
		timerStr = modelStyle.Render(fmt.Sprintf("%4s", m.timer.View()))
		spinnerStr = focusedModelStyle.Render(m.spinner.View())
	}

	if !m.altscreen {
		// Expense Tracker List
		header := listHeaderStyle.Render(fmt.Sprintf("%-20s | %-15s | %10s", "Title", "Category", "Value"))

		var rows []string

		if len(m.Expenses) == 0 {
			rows = append(rows, listItemStyle.Render("no expenses recorded yet."))
		}

		for _, e := range m.Expenses {
			var amountStr string
			val := FormatAmount(e.Spent)

			if e.Expense {
				amountStr = lossStyle.Render("-" + val)
			} else {
				amountStr = profitStyle.Render("+" + val)
			}

			catStr := dimmedStyle.Render(GetCategoryName(e.Category))

			row := fmt.Sprintf("%-20s | %-15s | %10s", e.Title, catStr, amountStr)
			rows = append(rows, listItemStyle.Render(row))
		}

		listContent := lipgloss.JoinVertical(lipgloss.Left, header, strings.Join(rows, "\n"))

		modeStatus := keywordStyle.Render(" List Mode ")
		status := fmt.Sprintf("\n\n You're in %s\n space: switch modes\n", modeStatus)

		return listContent + status
	} else {
		// Dasboard
		content = lipgloss.JoinHorizontal(lipgloss.Top, timerStr, spinnerStr)

		help := helpStyle.Render("\ntab: focus next * n: action * e: new expense * q: exit\n")
		modeStatus := keywordStyle.Render(" inline mode ")
		status := fmt.Sprintf("\n\n You're in %s\n space: switch modes\n", modeStatus)

		return content + help + status
	}
}
