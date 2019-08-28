package models

type RequestingUser struct {
	UserID      UserID
	Permissions map[string]bool
}

func (ru *RequestingUser) IsEmpty() bool {
	return ru.UserID == "" && len(ru.Permissions) == 0
}
