package searchresult

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/kaiquegarcia/gojooble/jooble"
)

type item jooble.Opportunity

func (i item) FilterValue() string {
	return ""
}

func itemsFromResponse(response *jooble.SearchResponse) []list.Item {
	output := make([]list.Item, len(response.Opportunities))
	for i, o := range response.Opportunities {
		output[i] = item(o)
	}

	return output
}
