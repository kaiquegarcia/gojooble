package searchresult

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/views/style"
	"github.com/kaiquegarcia/gojooble/jooble"
)

type (
	errMsg error
)
type model struct {
	list list.Model
	err  error
}

func newModel(
	response *jooble.SearchResponse,
) *model {
	list := list.New(
		itemsFromResponse(response),
		itemDelegate{},
		30,
		10,
	)
	list.Title = fmt.Sprintf("%d vagas encontradas:", len(response.Opportunities))
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.Styles.Title = style.TitleStyle
	list.Styles.PaginationStyle = style.PaginationStyle
	list.Styles.HelpStyle = style.HelpStyle

	return &model{
		list: list,
		err:  nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}
