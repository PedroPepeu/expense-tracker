// view.go
package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	// 1. Estados Especiais
	if m.suspending {
		return ""
	}
	if m.quitting {
		return "Bye!\n"
	}
	if m.FormState == StateForm {
		return m.form.View()
	}

	// 2. Preparação dos componentes do Dashboard (Timer/Spinner)
	// Calculamos isso antes para usar no 'else' lá embaixo
	var timerStr, spinnerStr string
	if m.state == TimerView {
		timerStr = focusedModelStyle.Render(fmt.Sprintf("%4s", m.timer.View()))
		spinnerStr = modelStyle.Render(m.spinner.View())
	} else {
		timerStr = modelStyle.Render(fmt.Sprintf("%4s", m.timer.View()))
		spinnerStr = focusedModelStyle.Render(m.spinner.View())
	}

	// 3. Decisão de Tela: Lista (AltScreen) vs Dashboard (Inline)
	if m.altscreen {
		// --- MODO LISTA DE DESPESAS (AltScreen) ---

		header := listHeaderStyle.Render(fmt.Sprintf("%-20s | %-15s | %10s", "Title", "Category", "Value"))

		var rows []string

		if len(m.Expenses) == 0 {
			rows = append(rows, listItemStyle.Render("No expenses recorded yet."))
		}

		for i, e := range m.Expenses {
			var amountStr string
			val := FormatAmount(e.Spent)

			// Cores: Vermelho para Despesa, Verde para Lucro
			if e.Expense {
				amountStr = lossStyle.Render("-" + val)
			} else {
				amountStr = profitStyle.Render("+" + val)
			}

			catStr := dimmedStyle.Render(GetCategoryName(e.Category))

			// Linha base sem formatação de seleção
			rowBase := fmt.Sprintf("%-20s | %-15s | %10s", e.Title, catStr, amountStr)

			// Lógica do Cursor: Verifica se o índice atual é o selecionado
			if i == m.Cursor {
				// Adiciona seta > e muda o estilo (selectedStyle precisa estar em styles.go)
				rows = append(rows, selectedStyle.Render("> "+rowBase))
			} else {
				// Adiciona espaços vazios para alinhar
				rows = append(rows, listItemStyle.Render("  "+rowBase))
			}
		}

		listContent := lipgloss.JoinVertical(lipgloss.Left, header, strings.Join(rows, "\n"))

		modeStatus := keywordStyle.Render(" List Mode ")
		// Adicionei as instruções de navegação no rodapé
		status := fmt.Sprintf("\n\n You're in %s\n arrows: nav * d: delete * space: switch modes * e: new expense\n", modeStatus)

		return listContent + status

	} else {
		// --- MODO DASHBOARD (Inline) ---

		content := lipgloss.JoinHorizontal(lipgloss.Top, timerStr, spinnerStr)

		help := helpStyle.Render("\ntab: focus next * n: action * space: switch modes * q: exit\n")
		modeStatus := keywordStyle.Render(" inline mode ")
		status := fmt.Sprintf("\n\n You're in %s\n", modeStatus)

		return content + help + status
	}
}
