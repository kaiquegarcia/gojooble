package setconfig

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

const inputsLength = 1

// inputs
const (
	apiKey = iota
)

// valid inputs counter to enable quit
var inputValidations = map[int]func(s string) bool{
	apiKey: func(s string) bool {
		return len(s) == 36
	},
}

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func (m model) Result() config.Configuration {
	return config.New(
		m.inputs[apiKey].Value(),
	)
}

func newModel(current config.Configuration) model {
	var inputs []textinput.Model = make([]textinput.Model, inputsLength)
	inputs[apiKey] = textinput.New()
	inputs[apiKey].Placeholder = "abcdefgh-1234-4567-abcd-abcdefghijkL"
	inputs[apiKey].Focus()
	inputs[apiKey].CharLimit = 36
	inputs[apiKey].Width = 40
	inputs[apiKey].Prompt = ""
	inputs[apiKey].SetValue(current.ApiKey())

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				shouldQuit := true
				for i, v := range m.inputs {
					if valid := inputValidations[i](v.Value()); !valid {
						m.focused = i
						shouldQuit = false
						break
					}
				}

				if shouldQuit {
					return m, tea.Quit
				}
			} else {
				m.nextInput()
			}
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
		` Change all configurations field until you end on the
 last input, then press ENTER to save. If any required field is
 not filled, the form won't be closed.

 Note: to generate a new API_KEY, access Jooble's form page:
 > https://jooble.org/api/about

 %s
 %s

 %s
`,
		style.InputStyle.Width(40).Render("API_KEY"),
		m.inputs[apiKey].View(),
		style.ButtonStyle.Render("Save"),
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
