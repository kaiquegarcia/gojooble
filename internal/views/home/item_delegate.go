package home

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gojooble/internal/views/style"
)

type itemDelegate struct{}

func (d itemDelegate) Height() int {
	return 1
}
func (d itemDelegate) Spacing() int {
	return 0
}
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	value, ok := listItem.(item)
	if !ok {
		return
	}

	styler := style.ItemStyle
	prefix := ""
	if index == m.Index() {
		styler = style.SelectedItemStyle
		prefix = "> "
	}

	str := fmt.Sprintf("%s%d. %s", prefix, index+1, value)
	fmt.Fprint(w, styler.Render(str))
}
