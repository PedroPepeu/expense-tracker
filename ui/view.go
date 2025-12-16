// view.go
package ui

import (
	"fmt"

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

	content = lipgloss.JoinVertical(lipgloss.Top, timerStr, spinnerStr)

	help := helpStyle.Render("\ntab: focus next * n: action * e: new expense * q: exit\n")

	var modeStatus string
	if m.altscreen {
		modeStatus = keywordStyle.Render(" altscreen mode ")
	} else {
		modeStatus = keywordStyle.Render(" inline mode ")
	}

	status := fmt.Sprintf("\n\n You're in %s\n space: switch modes\n", modeStatus)

	return content + help + status
}
