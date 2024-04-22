package screen

import (
	"fmt"
	"sort"

	"github.com/bbfh-dev/configure.mcvm/mcvm/auth"
	"github.com/bbfh-dev/configure.mcvm/mcvm/config"
	"github.com/bbfh-dev/configure.mcvm/ui/component"
	"github.com/bbfh-dev/configure.mcvm/ui/style"
	"github.com/bbfh-dev/configure.mcvm/ui/tools"
	"github.com/bbfh-dev/configure.mcvm/ui/widget"
	"github.com/bbfh-dev/configure.mcvm/ui/widget/field"
	tea "github.com/charmbracelet/bubbletea"
)

// ---- UserScreen

type UserScreen struct {
	title         string
	width         int
	height        int
	notifications []tools.Notification
	messages      []tools.Message
	editing       string

	userListWidget    widget.RadioWidget
	actionsWidget     widget.RadioWidget
	editUserWidget    widget.FormWidget
	editActionsWidget widget.RadioWidget
}

func NewUserScreen() UserScreen {
	return UserScreen{
		title:          tools.WithIcon(tools.USER_ICON, "Manage users"),
		userListWidget: widget.NewRadioWidget(NewUserList(), true),
		actionsWidget: widget.NewRadioWidget(component.NewLiteralList([]component.Item{
			component.NewLiteralItem(
				"new",
				tools.WithIcon(tools.NEW_USER_ICON, "New"),
				"Create a new user entry",
			),
			component.NewLiteralItem(
				"save",
				tools.WithIcon(tools.SAVE_ICON, "Save & Quit"),
				"Write the changes to disk",
			),
		}), false),
		editUserWidget: widget.NewFormWidget([]field.Field{
			field.NewField(
				"Id",
				"The identifier of the user",
				field.NewInputField(field.IdValidator{}),
			),
			field.NewField(
				"Type",
				"The type of authentication to use for the user",
				field.NewSelectField(component.NewChoiceList(&component.UserTypeOptions)),
			),
		}, true),
		editActionsWidget: widget.NewRadioWidget(component.NewLiteralList([]component.Item{
			component.NewLiteralItem(
				"set_default",
				tools.WithIcon(tools.STAR_ICON, "Set as default"),
				"Sets this user as the default one",
			),
			component.NewLiteralItem(
				"confirm",
				tools.WithIcon(tools.SAVE_ICON, "Confirm"),
				"Return to the user screen",
			),
			component.NewLiteralItem(
				"delete",
				tools.WithIcon(tools.DELETE_ICON, "Delete"),
				"Delete the user",
			),
		}), false),
		editing: "",
	}
}

func (screen UserScreen) Title() string {
	return screen.title
}

func (screen UserScreen) Update(raw tea.Msg) (Screen, tea.Cmd) {
	var commands []tea.Cmd

	switch msg := raw.(type) {

	case tea.WindowSizeMsg:
		screen.width = msg.Width
		screen.height = msg.Height
		screen.userListWidget.Width = msg.Width
		screen.actionsWidget.Width = msg.Width
	}

	if len(screen.editing) != 0 {
		screen.editUserWidget = tools.UpdateModel[widget.FormWidget](
			&commands,
			screen.editUserWidget.Clear(),
			raw,
		)
		screen.editActionsWidget = tools.UpdateModel[widget.RadioWidget](
			&commands,
			screen.editActionsWidget.Clear(),
			raw,
		)
		// Messages
		for _, message := range screen.editUserWidget.Messages {
			switch message.(type) {
			case tools.OverflowBottom:
				screen.editUserWidget = screen.editUserWidget.Blur()
				screen.editActionsWidget = screen.editActionsWidget.Focus()
				screen.editActionsWidget.Cursor = 0
				return screen.Update(nil)
			}
		}
		for _, message := range screen.editActionsWidget.Messages {
			switch msg := message.(type) {
			case tools.OverflowTop:
				screen.editUserWidget = screen.editUserWidget.Focus()
				screen.editActionsWidget = screen.editActionsWidget.Blur()
				return screen.Update(nil)

			case component.SelectItem:
				switch msg.Item {
				case "set_default":
					config.MCVMConfig.DefaultUser = screen.editing
					screen.editing = ""
				case "confirm":
					id := screen.editUserWidget.Fields[0].Type.Id()
					if len(id) > 0 {
						user := config.MCVMConfig.Users[screen.editing]
						delete(config.MCVMConfig.Users, screen.editing)
						user.Type = screen.editUserWidget.Fields[1].Type.Id()
						config.MCVMConfig.Users[id] = user
					}
					screen.editing = ""
				case "delete":
					delete(config.MCVMConfig.Users, screen.editing)
					screen.userListWidget.Cursor = 0
					screen.editing = ""
				}
			}
		}

		screen.notifications = append(screen.notifications, screen.editUserWidget.Notification)
		screen.notifications = append(screen.notifications, screen.editActionsWidget.Notification)
	} else {
		screen.userListWidget = tools.UpdateModel[widget.RadioWidget](
			&commands,
			screen.userListWidget.Clear(),
			raw,
		)
		screen.actionsWidget = tools.UpdateModel[widget.RadioWidget](
			&commands,
			screen.actionsWidget.Clear(),
			raw,
		)

		// Messages
		for _, message := range screen.userListWidget.Messages {
			switch msg := message.(type) {
			case tools.OverflowBottom:
				screen.userListWidget = screen.userListWidget.Blur()
				screen.actionsWidget = screen.actionsWidget.Focus()
				screen.actionsWidget.Cursor = 0
				return screen.Update(nil)

			case component.SelectItem:
				screen.editing = msg.Item
				screen.editUserWidget.Fields[0].Type = screen.editUserWidget.Fields[0].Type.Set(msg.Item)
			}
		}
		for _, message := range screen.actionsWidget.Messages {
			switch msg := message.(type) {
			case tools.OverflowTop:
				screen.userListWidget = screen.userListWidget.Focus()
				screen.actionsWidget = screen.actionsWidget.Blur()
				screen.userListWidget.Cursor = screen.userListWidget.List.Size() - 1
				return screen.Update(nil)

			case component.SelectItem:
				switch msg.Item {
				case "new":
					screen.editing = "unnamed"
					screen.editUserWidget.Fields[0].Type = screen.editUserWidget.Fields[0].Type.Set("")
					screen.editUserWidget.Fields[1].Type = screen.editUserWidget.Fields[1].Type.Set("microsoft")
				case "save":
					screen.messages = append(screen.messages, widget.SaveMsg{})
				}
			}
		}

		screen.notifications = append(screen.notifications, screen.userListWidget.Notification)
		screen.notifications = append(screen.notifications, screen.actionsWidget.Notification)
		screen.notifications = append(screen.notifications, screen.editUserWidget.Notification)
	}

	return screen, tea.Batch(commands...)
}

func (screen UserScreen) View(width int) []string {
	var contents []string

	if len(screen.editing) == 0 {
		tools.AppendContent(&contents, style.Text.Render(width, "User list:"), width)
		tools.AppendContent(&contents, screen.userListWidget.View(), width)
		tools.AppendContent(&contents, screen.actionsWidget.View(), width)
	} else {
		tools.AppendContent(&contents, screen.editUserWidget.View(), width)
		tools.AppendContent(&contents, screen.editActionsWidget.View(), width)
	}

	return contents
}

func (screen UserScreen) Clear() Screen {
	screen.notifications = make([]tools.Notification, 0)
	screen.messages = make([]tools.Message, 0)

	return screen
}

func (screen UserScreen) Notifications() []tools.Notification {
	return screen.notifications
}

func (screen UserScreen) Messages() []tools.Message {
	return screen.messages
}

func (screen UserScreen) Keys() tools.Keybinds {
	return screen.userListWidget.Keys
}

func (screen UserScreen) Lock() bool {
	return len(screen.editing) != 0 && screen.editUserWidget.Editing
}

// ---- List

type UserList struct{}

func NewUserList() UserList {
	return UserList{}
}

func (list UserList) Items() []component.Item {
	var items []component.Item

	var keys []string
	for key := range config.MCVMConfig.Users {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		items = append(items, UserItem{key})
	}
	return items
}

func (list UserList) Size() int {
	return len(config.MCVMConfig.Users)
}

// ---- Item

type UserItem struct {
	id string
}

func NewUserItem(id string) UserItem {
	return UserItem{
		id: id,
	}
}

func (item UserItem) Id() string {
	return item.id
}

func (item UserItem) Title() string {
	isDefault := ""
	if config.MCVMConfig.DefaultUser == item.id {
		isDefault = fmt.Sprintf(" [%s]", tools.WithIcon(tools.STAR_ICON, "Default"))
	}

	user, ok := auth.MCVMAuth.Users[item.id]
	if !ok {
		return fmt.Sprintf("%s*%s (Not authenticated)", item.id, isDefault)
	}

	return fmt.Sprintf("%s%s", user.Username, isDefault)
}

func (item UserItem) Description() string {
	return item.id
}
