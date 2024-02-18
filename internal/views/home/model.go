package home

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/views/style"
)

type (
	errMsg error
)

// options
const (
	OptSearch    item = "Search Jobs"
	OptSetconfig item = "Change configuration"
	optQuit      item = "Quit"
)

type model struct {
	list   list.Model
	choice item
	err    error
}

func newModel() *model {
	list := list.New(
		[]list.Item{
			OptSearch,
			OptSetconfig,
			optQuit,
		},
		itemDelegate{},
		30,
		10,
	)
	list.Title = "What do you wanna do?"
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.Styles.Title = style.TitleStyle
	list.Styles.PaginationStyle = style.PaginationStyle
	list.Styles.HelpStyle = style.HelpStyle

	return &model{
		list:   list,
		choice: "",
		err:    nil,
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
			value, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = value
			}

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
