package setconfig

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/config"
)

func Run(current config.Configuration) (config.Configuration, error) {
	model := newModel(current)
	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gojooble > config")

	if _, err := program.Run(); err != nil {
		return nil, err
	}

	c := model.Result()
	return c, config.Save(c)
}
