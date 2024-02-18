package loading

import (
	tea "github.com/charmbracelet/bubbletea"
)

var m *model

func Start() {
	if m != nil {
		return
	}

	m = newModel()
	go func() {
		program := tea.NewProgram(m, tea.WithAltScreen())
		program.SetWindowTitle("#gojooble > carregando...")
		program.Run()
		m.quitSignal <- true
	}()
}

func Stop() {
	if m == nil {
		return
	}

	m.quit = true
	<-m.quitSignal
	m = nil
}
