package searchresult

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
	opportunity, ok := listItem.(item)
	if !ok {
		return
	}

	styler := style.JobItemStyle
	prefix := " "
	if index == m.Index() {
		styler = style.SelectedItemStyle
		prefix = ">"
	}

	str := fmt.Sprintf(
		`%s %d. [%s] %s (%v)
    ~ Link:  %s`,
		prefix, index+1, opportunity.Location, opportunity.Title, opportunity.ID,
		opportunity.Link,
	)
	fmt.Fprint(w, styler.Render(str))
}
