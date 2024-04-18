package user_screen

import (
	"fmt"

	"github.com/bbfh-dev/configure.mcvm/mcvm"
	"github.com/bbfh-dev/configure.mcvm/mcvm/schema"
)

type UserItem struct {
	id   string
	user *schema.User
}

func NewUserItem(id string, user *schema.User) UserItem {
	return UserItem{
		id:   id,
		user: user,
	}
}

func (item UserItem) Id() string {
	return item.id
}

func (item UserItem) Name() string {
	if item.id == mcvm.MCVMConfig.DefaultUser {
		return fmt.Sprintf("%s (Default)", item.id)
	}

	return item.id
}

func (item UserItem) Description() string {
	if _, ok := mcvm.MCVMAuth.Users[item.id]; !ok {
		return fmt.Sprintf("%s - (Not authenticated)", *item.user.Type)
	}

	return fmt.Sprintf("%s - %s", *item.user.Type, mcvm.MCVMAuth.Users[item.id].UUID)
}
