package search

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/config"
	"github.com/kaiquegarcia/gojooble/internal/views/style"
)

type (
	errMsg error
)

const inputsLength = 2

// inputs
const (
	keywords = iota
	location
)

type model struct {
	inputs  []textinput.Model
	inputed bool
	focused int
	err     error
}

type result struct {
	keywords string
	location string
}

func (m model) Result() result {
	return result{
		keywords: m.inputs[keywords].Value(),
		location: m.inputs[location].Value(),
	}
}

func newModel(current config.Configuration) *model {
	var inputs []textinput.Model = make([]textinput.Model, inputsLength)
	inputs[keywords] = textinput.New()
	inputs[keywords].Placeholder = "golang"
	inputs[keywords].Focus()
	inputs[keywords].CharLimit = 30
	inputs[keywords].Width = 34
	inputs[keywords].Prompt = ""

	inputs[location] = textinput.New()
	inputs[location].Placeholder = "Alabama (optional)"
	inputs[location].CharLimit = 30
	inputs[location].Width = 34
	inputs[location].Prompt = ""

	return &model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				m.inputed = true
				return m, tea.Quit
			}

			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		` Now it's time to look up for a job!
 Fill the fields above and press ENTER to search.

 %s  %s
 %s %s

 %s
`,
		style.InputStyle.Width(34).Render("Keywords"),
		style.InputStyle.Width(34).Render("Location"),
		m.inputs[keywords].View(),
		m.inputs[location].View(),
		style.ButtonStyle.Render("Search"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
