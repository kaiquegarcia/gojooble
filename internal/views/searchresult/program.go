package searchresult

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/jooble"
)

func Run(
	response *jooble.SearchResponse,
) error {
	program := tea.NewProgram(newModel(response), tea.WithAltScreen())
	program.SetWindowTitle("#gojooble > search > result")
	_, err := program.Run()
	return err
}
