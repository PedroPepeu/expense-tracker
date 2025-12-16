// models.go
package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/huh"
)

// Consts

const (
	DefaultTime = time.Minute

	// View States
	TimerView sessionState = iota
	SpinnerView

	// App States
	StateList = iota
	StateForm
)

// Types
type sessionState uint

// Main UI Model
type Expense struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Category    int       `json:"category"`
	Spent       int       `json:"spent"`
	Installment int       `json:"installment"`
	Expense     bool      `json:"expense"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created"`
}

// Form Data Helper
type FormData struct {
	Title       string
	Category    int
	Spent       string
	Installment string
	IsExpense   int
	Date        string
	Confirm     bool
}

func NewExpenseForm(data *FormData) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What where the expense?").
				Prompt(":").
				Placeholder("Lunch").
				Validate(ValidateName).
				Value(&data.Title),

			huh.NewSelect[int]().
				Title("What was the category of the expense?").
				Options(
					huh.NewOption("Fun", 0),
					huh.NewOption("Work", 1),
					huh.NewOption("Study", 2),
					huh.NewOption("Investment", 3),
				).
				Value(&data.Category),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("How much do you Spent?").
				Prompt("$").
				Placeholder("30,99").
				Validate(ValidateAmount).
				Value(&data.Spent),

			huh.NewInput().
				Title("How many times?").
				Prompt("x").
				Placeholder("5").
				Validate(ValidateInt).
				Value(&data.Installment),
		),
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Type?").
				Options(
					huh.NewOption("Profit", 0),
					huh.NewOption("Loss", 1),
				).
				Value(&data.IsExpense),

			huh.NewInput().
				Title("When it was?").
				Prompt(":").
				Placeholder("DD-MM-YYYY").
				// Validate(ValidateDate).
				Value(&data.Date),

			huh.NewConfirm().
				Title("Do you want to save?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&data.Confirm),
		),
	).WithTheme(huh.ThemeBase())
}

type Model struct {
	// Expenses
	Expenses []Expense

	// Composable
	state   sessionState
	timer   timer.Model
	spinner spinner.Model
	index   int

	// Altscreen model
	altscreen  bool
	quitting   bool
	suspending bool

	// Form
	FormState   int
	form        *huh.Form
	formData    *FormData
	indexToEdit int

	Cursor int
}
