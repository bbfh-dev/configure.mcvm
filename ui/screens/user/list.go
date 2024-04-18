package user_screen

import (
	"sort"

	"github.com/bbfh-dev/configure.mcvm/mcvm"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
)

type UserList struct {
}

func NewUserList(items ...UserItem) UserList {
	return UserList{}
}

func (list UserList) Items() []tools.ListItem {
	var items []tools.ListItem
	for key, user := range mcvm.MCVMConfig.Users {
		items = append(items, NewUserItem(key, &user))
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Id() < items[j].Id()
	})

	return append(
		items,
		tools.NewLiteralItem("new", tools.WithIcon(tools.NEW_USER_ICON, "New"), "Add new user"),
		tools.NewLiteralItem(
			"save",
			tools.WithIcon(tools.SAVE_ICON, "Save"),
			"Writes the changes to disk",
		),
	)
}
