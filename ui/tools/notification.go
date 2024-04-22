package tools

type NotificationType uint8

const (
	HELP NotificationType = iota
	TASK
	ERROR
)

type Notification struct {
	Type NotificationType
	Text string
}
