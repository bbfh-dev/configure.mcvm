package component

type Item interface {
	Id() string
	Title() string
	Description() string
}

type LiteralItem struct {
	id          string
	title       string
	description string
}

func NewLiteralItem(id string, title string, description string) LiteralItem {
	return LiteralItem{
		id:          id,
		title:       title,
		description: description,
	}
}

func (item LiteralItem) Id() string {
	return item.id
}

func (item LiteralItem) Title() string {
	return item.title
}

func (item LiteralItem) Description() string {
	return item.description
}
