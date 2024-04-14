package schema

type User struct {
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}
