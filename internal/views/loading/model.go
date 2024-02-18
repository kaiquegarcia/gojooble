package loading

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/views/style"
)

type (
	errMsg error
)

type model struct {
	spinner    spinner.Model
	quit       bool
	quitSignal chan bool
	err        error
}

func newModel() *model {
	return &model{
		spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(style.SpinnerStyle),
		),
		quitSignal: make(chan bool),
		err:        nil,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.quit {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.spinner.View()
}
