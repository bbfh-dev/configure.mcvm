package tools

type ListItem interface {
	Id() string
	Name() string
	Description() string
}

type LiteralItem struct {
	id          string
	name        string
	description string
}

func NewLiteralItem(id string, name string, description string) LiteralItem {
	return LiteralItem{
		id:          id,
		name:        name,
		description: description,
	}
}

func (item LiteralItem) Id() string {
	return item.id
}

func (item LiteralItem) Name() string {
	return item.name
}

func (item LiteralItem) Description() string {
	return item.description
}
