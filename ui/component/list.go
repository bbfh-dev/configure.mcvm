package component

type SelectItem struct {
	Item string
}

type List interface {
	Items() []Item
	Size() int
}

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
