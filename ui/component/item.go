package component

// A generic list item
type Item interface {
	// Used to identify the item within the program
	Id() string

	// How the item will be displayed
	Title() string

	// Help notification when the item is hovered
	Description() string
}

// A list item with static values
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
