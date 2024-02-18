package search

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/config"
	"github.com/kaiquegarcia/gojooble/internal/views/loading"
	"github.com/kaiquegarcia/gojooble/jooble"
)

func Run(
	conf config.Configuration,
	joobleClient jooble.Jooble,
) (*jooble.SearchResponse, error) {
	model := newModel(conf)
	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gojooble > search")

	if _, err := program.Run(); err != nil {
		return nil, err
	}

	if !model.inputed {
		return nil, nil
	}

	r := model.Result()

	loading.Start()
	defer loading.Stop()

	response, _, err := joobleClient.Search(r.keywords, r.location)
	return response, err
}
