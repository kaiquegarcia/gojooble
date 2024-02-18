package style

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	InputStyle        = lipgloss.NewStyle().Foreground(LightBlue)
	ButtonStyle       = lipgloss.NewStyle().Foreground(DarkGray)
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	JobItemStyle      = lipgloss.NewStyle().PaddingLeft(2)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(BlueGreen)
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	SpinnerStyle      = lipgloss.NewStyle().Foreground(BlueGreen)
)
