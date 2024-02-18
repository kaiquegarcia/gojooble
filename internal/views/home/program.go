package home

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() (item, error) {
	model := newModel()
	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gojooble")
	if _, err := program.Run(); err != nil {
		return "", err
	}

	return model.choice, nil
}
