package component

type SelectItemMsg struct {
	Item string
}

// A generic list
type List interface {
	// Returns items to be iterated over
	Items() []Item

	// The size of the list
	Size() int
}

// A static list
type LiteralList struct {
	items []Item
}

func NewLiteralList(items []Item) LiteralList {
	return LiteralList{
		items: items,
	}
}

func (list LiteralList) Items() []Item {
	return list.items
}

func (list LiteralList) Size() int {
	return len(list.items)
}
