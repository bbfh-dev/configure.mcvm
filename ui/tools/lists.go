package tools

type List interface {
	Items() []ListItem
}

type LiteralList struct {
	items []ListItem
}

func NewLiteralList(items ...ListItem) LiteralList {
	return LiteralList{
		items: items,
	}
}

func (list LiteralList) Items() []ListItem {
	return list.items
}
